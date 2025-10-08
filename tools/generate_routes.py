import re
import sys
from pathlib import Path
from typing import Any, Dict, List, Optional, Tuple
import yaml
from jinja2 import Environment, FileSystemLoader

ROOT = Path(__file__).resolve().parents[1]
TEMPLATES = ROOT / "tools" / "templates"
ROUTES_DIR = ROOT / "app" / "routes"

env = Environment(loader=FileSystemLoader(str(TEMPLATES)), autoescape=False)

# Map (type, format) -> Python types
PY_TYPE_MAP: Dict[Tuple[Optional[str], Optional[str]], str] = {
    ("string", None): "str",
    ("string", "uuid"): "UUID",
    ("string", "date"): "date",
    ("string", "date-time"): "datetime",
    ("integer", None): "int",
    ("number", None): "float",
    ("boolean", None): "bool",
}

HTTP_METHODS = ["get", "post", "put", "patch", "delete", "options", "head"]


def sanitize_module(s: str) -> str:
    # Keep your names, only make them filesystem-safe: spaces/colons -> underscore
    s = s.strip().lower().replace(":", "_")
    s = re.sub(r"[^0-9a-z_]+", "_", s)
    return s or "default"


def sanitize_func_name(s: str) -> str:
    return re.sub(r"[^0-9a-zA-Z_]", "_", s)


def default_func_name(method: str, path: str) -> str:
    return f"{method.lower()}_{re.sub(r'[^0-9a-zA-Z_]', '_', path.strip('/')) or 'root'}"


def ref_name(ref: str) -> str:
    return ref.split("/")[-1]


def type_hint_for_schema(schema: Dict[str, Any]) -> str:
    if not schema:
        return "Any"
    if "$ref" in schema:
        return ref_name(schema["$ref"])
    t = schema.get("type")
    fmt = schema.get("format")
    # arrays
    if t == "array":
        items = schema.get("items", {}) or {}
        inner = type_hint_for_schema(items)
        return f"List[{inner}]"
    # primitives
    return PY_TYPE_MAP.get((t, fmt)) or PY_TYPE_MAP.get((t, None)) or "Any"


def type_hint_for_param(p: Dict[str, Any]) -> str:
    return type_hint_for_schema(p.get("schema", {}))


def collect_params(path_params: List[Dict[str, Any]], op_params: List[Dict[str, Any]]) -> List[Dict[str, Any]]:
    merged: Dict[Tuple[str, str], Dict[str, Any]] = {}
    for p in (path_params or []) + (op_params or []):
        key = (str(p.get("name") or ""), str(p.get("in") or ""))
        merged[key] = p
    params = []
    for _, p in merged.items():
        hint = type_hint_for_param(p)
        default = None if p.get("required") else "None"
        params.append({
            "name": p["name"],
            "type_hint": hint,
            "default": default,
            "in": p.get("in"),
        })
    return params


def detect_paginated_response(op: Dict[str, Any]) -> Optional[str]:
    """
    Detect your pattern:
    responses:
      200:
        content:
          application/json:
            schema:
              allOf:
                - $ref: '#/components/schemas/PaginatedResponse'
                - { type: object, properties: { results: { items: { $ref: ... }}}}
    Returns the item model name if found, else None.
    """
    responses = op.get("responses", {}) or {}
    cand = None
    for code in ["200", "201", "204"]:
        if code in responses:
            cand = responses[code]
            break
    if not cand:
        if responses:
            cand = next(iter(responses.values()))
        else:
            return None
    content = (cand or {}).get("content", {})
    media = content.get(
        "application/json") or (next(iter(content.values()), None) if content else None)
    if not media:
        return None
    schema = media.get("schema") or {}
    if not isinstance(schema, dict):
        return None
    all_of = schema.get("allOf")
    if not isinstance(all_of, list):
        return None
    # look for PaginatedResponse + results array $ref
    saw_paginated = any(("$ref" in part and ref_name(
        part["$ref"]) == "PaginatedResponse") for part in all_of if isinstance(part, dict))
    if not saw_paginated:
        return None
    for part in all_of:
        if not isinstance(part, dict):
            continue
        props = (part.get("properties") or {}) if part.get(
            "type") == "object" else {}
        results = props.get("results")
        items = (results or {}).get(
            "items") if isinstance(results, dict) else None
        if isinstance(items, dict) and "$ref" in items:
            return ref_name(items["$ref"])
    return None


def first_response_model(op: Dict[str, Any]) -> Tuple[Optional[str], Optional[int], Optional[str]]:
    """
    Standard model detection; plus returns 'paginated_item' if your paginated pattern is detected.
    """
    responses = op.get("responses", {}) or {}
    chosen_code = None
    for code in ["200", "201", "204"]:
        if code in responses:
            chosen_code = code
            break
    if not chosen_code and responses:
        chosen_code = next(iter(responses.keys()))
    if not chosen_code:
        return None, None, None

    model, status = None, int(chosen_code) if chosen_code.isdigit() else None
    content = responses[chosen_code].get(
        "content", {}) if responses.get(chosen_code) else {}
    media = content.get(
        "application/json") or (next(iter(content.values()), None) if content else None)
    if media:
        schema = media.get("schema", {})
        if "$ref" in schema:
            model = ref_name(schema["$ref"])
        elif schema.get("type") == "array" and "$ref" in (schema.get("items") or {}):
            model = f"List[{ref_name(schema['items']['$ref'])}]"
        # leave others as None (we’ll hint 'dict' or 'Any')
    paginated_item = detect_paginated_response(op)
    return model, status, paginated_item


def request_body_model(op: Dict[str, Any]) -> Optional[str]:
    rb = op.get("requestBody", {})
    content = rb.get("content", {}) if isinstance(rb, dict) else {}
    if not content:
        return None
    media = content.get(
        "application/json") or next(iter(content.values()), None)
    if not media:
        return None
    schema = media.get("schema", {})
    if "$ref" in schema:
        return ref_name(schema["$ref"])
    if schema.get("type") == "array" and "$ref" in (schema.get("items") or {}):
        return f"List[{ref_name(schema['items']['$ref'])}]"
    return "Any"


def order_tags(spec: Dict[str, Any], discovered: List[str]) -> List[str]:
    ordered: List[str] = []
    seen = set()
    for group in spec.get("x-tagGroups", []) or []:
        for t in group.get("tags", []) or []:
            if t in discovered and t not in seen:
                ordered.append(t)
                seen.add(t)
    for t in discovered:
        if t not in seen:
            ordered.append(t)
    return ordered


def main(spec_path: str = "openapi.yaml") -> None:
    with open(spec_path, "r", encoding="utf-8") as f:
        spec = yaml.safe_load(f)

    paths = spec.get("paths", {}) or {}
    groups: Dict[str, List[Dict[str, Any]]] = {}
    discovered_tags: List[str] = []

    for path, path_item in paths.items():
        path_params = (path_item or {}).get("parameters", []) or []
        for method in HTTP_METHODS:
            if method not in path_item:
                continue
            op = path_item[method]
            tags = op.get("tags") or ["default"]
            public_tag = tags[0]
            discovered_tags.append(public_tag)

            op_params = collect_params(
                path_params, op.get("parameters", []) or [])
            resp_model, status_code, paginated_item = first_response_model(op)
            body_model = request_body_model(op)

            func_name = sanitize_func_name(
                op.get("operationId") or default_func_name(method, path))

            operations = groups.setdefault(public_tag, [])
            notes = "Auto-generated stub. Add auth/deps if needed."
            if paginated_item:
                notes += f" PaginatedResponse detected (results: List[{paginated_item}])."

            operations.append({
                "method": method,
                "path": path,
                "func_name": func_name,
                "operation_id": op.get("operationId", ""),
                "summary": op.get("summary"),
                "params": op_params,
                "body": None if body_model is None else {"model": body_model},
                # leave None for complex/allOf to avoid wrong typing
                "response_model": resp_model,
                "status_code": status_code,
                "return_hint": "dict" if paginated_item else (resp_model or "None"),
                "notes": notes,
            })

    # Render per-tag modules
    ROUTES_DIR.mkdir(parents=True, exist_ok=True)
    tmplt = env.get_template("route_module.j2")

    for public_tag, operations in groups.items():
        module = sanitize_module(public_tag)
        out_file = ROUTES_DIR / f"{module}.py"
        code = tmplt.render(
            public_tag=public_tag,
            module_stub=module,
            operations=sorted(operations, key=lambda x: x["func_name"])
        )
        out_file.write_text(code, encoding="utf-8")

    # Register routers in x-tagGroups order
    ordered = order_tags(spec, [t for t in discovered_tags])
    uniq_ordered = []
    seen = set()
    for t in ordered:
        if t not in seen:
            uniq_ordered.append(t)
            seen.add(t)

    imports_tpl = env.get_template("router_imports.j2")
    tags_render = [
        {"public": t, "module": sanitize_module(t)} for t in uniq_ordered]
    (ROOT / "app" / "routes" / "_register.py").write_text(
        imports_tpl.render(tags=tags_render),
        encoding="utf-8"
    )

    print("Generated route modules:", ", ".join(
        [sanitize_module(t) for t in uniq_ordered]))
    print("Wrote router register: app/routes/_register.py")


if __name__ == "__main__":
    spec = sys.argv[1] if len(sys.argv) > 1 else "openapi.yaml"
    main(spec)
