# FastAPI Codegen Pipeline — Spec-First Workflow

This guide describes how to generate, validate, and maintain a **FastAPI server automatically** from your OpenAPI specification. It covers:

1. One-time bootstrapping using **OpenAPI Generator**.
2. A **reproducible pipeline** for regenerating Pydantic models and route stubs whenever your spec changes.
3. Practical fixes for known generator bugs and environment quirks.

## 1. Prerequisites (One-Time Setup)

### Required Tools

| Tool                  | Purpose                                       |
| --------------------- | --------------------------------------------- |
| **Python 3.11+**      | Runtime and FastAPI environment               |
| **Node.js**           | Simplest way to install OpenAPI Generator CLI |
| **pipx** *(optional)* | Isolated CLI installs                         |

### Install Dependencies

```bash
# OpenAPI Generator CLI
npm install -g @openapitools/openapi-generator-cli

# Datamodel Code Generator (Pydantic models)
pipx install datamodel-code-generator || pip install datamodel-code-generator

# Validation & templating utilities
pipx install openapi-spec-validator || pip install openapi-spec-validator
pip install pyyaml jinja2 black isort
```

### Validate the Spec

```bash
openapi-spec-validator docs/api/openapi.yaml
```

Fix any reported errors before continuing.

## 2. One-Time Bootstrap — Generate a Runnable Server

This step verifies that your OpenAPI spec can be implemented.

```bash
rm -rf generated_fastapi
openapi-generator-cli generate \
  -i docs/api/openapi.yaml \
  -g python-fastapi \
  -o ./generated_fastapi
```

### Post-Generation Patches

Fix known template bugs from the official generator:

```bash
# Quote missing string defaults
find generated_fastapi/src/openapi_server/apis -type f -name "*.py" \
  -exec sed -i 's/= Query(outgoing, /= Query("outgoing", /g' {} +

# Clean up invalid HTML escapes
find generated_fastapi/src/openapi_server/apis -type f -name "*.py" \
  -exec sed -i 's/\\&/&/g' {} +
```

### Run the Generated Server

Install dependencies:

```bash
uv pip install -r generated_fastapi/requirements.txt
```

Run from your **project root**:

```bash
PYTHONPATH=generated_fastapi/src uv run uvicorn openapi_server.main:app \
  --reload --reload-exclude .venv --port 8080
```

Open [http://127.0.0.1:8080/docs](http://127.0.0.1:8080/docs) to verify your routes appear.

> **Tip:** Don’t commit `generated_fastapi/`.
> It’s a disposable scaffold used only to confirm the spec works.

## 3. Establish a Real Project Structure

Copy only what you need from the scaffold into your actual repo:

```txt
your_api/
├─ app/
│  ├─ main.py
│  ├─ core/            # auth, settings, deps
│  ├─ schemas/         # generated Pydantic models
│  ├─ routes/          # generated route stubs
│  ├─ services/        # hand-written business logic
│  ├─ db/              # data-access layer
│  └─ utils/
├─ tools/
│  ├─ generate_routes.py
│  └─ templates/
│     ├─ route_module.j2
│     └─ router_imports.j2
├─ docs/api/openapi.yaml
├─ requirements.txt
└─ pyproject.toml
```

## 3️4. Generate Pydantic Models

```bash
datamodel-codegen \
  --input docs/api/openapi.yaml \
  --input-file-type openapi \
  --output app/schemas/models.py \
  --target-python-version 3.12 \
  --use-schema-description \
  --reuse-model \
  --strict-nullable
```

> **If you see:**
> `ImportError: cannot import name 'ParameterSource' from 'click.core'`
> → run `uv pip install --upgrade black click` (ensure `black>=24.3.0`, `click>=8.1.7`).

## 5. Generate Route Stubs with Jinja2

### 5.1 `tools/templates/route_module.j2`

```jinja
{% set module_doc = "Auto-generated from OpenAPI. Edits may be overwritten. Put business logic in app/services." %}
# {{ module_doc }}
from typing import Any, Optional, List, Dict
from uuid import UUID
from datetime import date, datetime
from fastapi import APIRouter, Depends, Path, Query, Body
from app.schemas.models import *  # noqa

router = APIRouter(tags=["{{ public_tag }}"])

{% for op in operations %}
@router.{{ op.method }}(
    "{{ op.path }}",
    {% if op.response_model %}response_model={{ op.response_model }},{% endif %}
    {% if op.status_code %}status_code={{ op.status_code }},{% endif %}
)
async def {{ op.func_name }}(
    {%- for p in op.params -%}
    {{ p.name }}: {{ p.type_hint }}{{ " = " + p.default if p.default is not none else "" }},
    {%- endfor -%}
    {%- if op.body is not none -%}
    body: {{ op.body.model }} = Body(...),
    {%- endif -%}
) -> {{ op.return_hint }}:
    """
    {{ op.summary or "" }}
    OperationId: {{ op.operation_id }}
    {{ op.notes }}
    """
    # TODO: Implement business logic in app/services.
    raise NotImplementedError("{{ op.method|upper }} {{ op.path }} not implemented yet")
{% endfor %}
```

### 5.2 `tools/templates/router_imports.j2`

```jinja
# Auto-generated router imports (ordered by x-tagGroups)
from fastapi import FastAPI
{% for tag in tags %}
from app.routes.{{ tag.module }} import router as {{ tag.module }}_router
{% endfor %}

def register_routers(app: FastAPI) -> None:
    {% for tag in tags %}
    app.include_router({{ tag.module }}_router)
    {% endfor %}
```

### 5.3 `tools/generate_routes.py`

```python
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
        if "$ref" in p:
            ref = p["$ref"].split("/")[-1]
            merged[(ref, "ref")] = {"name": ref,
                                    "in": "ref", "schema": {"type": "Any"}}
            continue

        key = (str(p.get("name") or ""), str(p.get("in") or ""))
        merged[key] = p

    params = []
    for _, p in merged.items():
        name = p.get("name")
        if not name:
            continue
        hint = type_hint_for_param(p)
        default = None if p.get("required") else "None"
        params.append({
            "name": name,
            "type_hint": hint,
            "default": default,
            "in": p.get("in"),
        })

    params.sort(key=lambda p: p["default"] is not None)
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
```

The generator parses your spec, groups endpoints by tag, and writes route modules.
Your fixed version includes:

* `$ref` handling for parameters
* required-parameter sorting (prevents syntax errors)
* safe type hints

### Run it

```bash
python tools/generate_routes.py docs/api/openapi.yaml
```

It creates one file per tag and a `app/routes/_register.py`.

## 6. Minimal Application Entrypoint

```python
# app/main.py
from fastapi import FastAPI
from app.routes._register import register_routers

def create_app() -> FastAPI:
    app = FastAPI(title="Your API", version="1.0.0")
    register_routers(app)
    return app

app = create_app()
```

### Run it

```bash
uv run uvicorn app.main:app --reload --port 8080
```

## 7. Common Fixes (Known Issues Solved)

| Symptom                                                        | Cause                            | Fix                                                |
| -------------------------------------------------------------- | -------------------------------- | -------------------------------------------------- |
| `NameError: outgoing not defined`                              | Generator bug                    | Quoted with sed patch above                        |
| `invalid escape sequence '\&'`                                 | Escaped HTML                     | Optional cleanup sed patch                         |
| `FileNotFoundError ... .venv/bin/python3.13`                   | Watcher scanning wrong venv      | Use `--reload-exclude .venv` or run without reload |
| `KeyError: 'name'`                                             | `$ref` parameters missing `name` | Skip or dereference `$ref` in `collect_params`     |
| `parameter without a default follows parameter with a default` | Required/optional out of order   | Sort params before returning                       |
| `_register.py` invalid syntax                                  | Wrong template key               | Use `{{ tag.module }}`                             |
| Port already in use                                            | Another Uvicorn running          | `pkill -f uvicorn` or change `--port`              |

## 8. Regeneration Workflow (No Drift)

Re-run any time your spec changes:

```bash
# Clean and regenerate
rm -rf generated_fastapi
openapi-generator-cli generate \
  -i docs/api/openapi.yaml \
  -g python-fastapi \
  -o ./generated_fastapi

# Apply small patches (if needed)
find generated_fastapi/src/openapi_server/apis -type f -name "*.py" \
  -exec sed -i 's/= Query(outgoing, /= Query("outgoing", /g' {} +
find generated_fastapi/src/openapi_server/apis -type f -name "*.py" \
  -exec sed -i 's/\\&/&/g' {} +

# Generate schemas and routes
datamodel-codegen --input docs/api/openapi.yaml --input-file-type openapi \
  --output app/schemas/models.py --target-python-version 3.12 --reuse-model --strict-nullable

python tools/generate_routes.py docs/api/openapi.yaml
```

## 9. Security, Dependencies, and Business Logic

* Add `app/core/security.py` for auth dependencies (`Depends(get_current_user)`).
* Centralize shared dependencies (DB sessions, pagination) in `app/core/deps.py`.
* Implement business logic inside `app/services/` — **never** edit generated routes directly.

## 10. Testing Contract Compliance

Use **Schemathesis** to verify that your FastAPI app conforms to the OpenAPI spec:

```bash
pip install schemathesis
schemathesis run --checks all http://127.0.0.1:8080/openapi.json
```

## 11. Optional: Containerization

`Dockerfile`

```dockerfile
FROM python:3.12-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY app ./app
COPY docs/api/openapi.yaml ./openapi.yaml
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]
```

## Version Control Guidelines

| File / Folder                 | Commit? | Notes                                        |
| ----------------------------- | ------- | -------------------------------------------- |
| `docs/api/openapi.yaml`       | ✅       | Source of truth                              |
| `openapitools.json`           | ✅       | Pins generator version and config            |
| `generated_fastapi/`          | 🚫      | Disposable output                            |
| `app/schemas/`, `app/routes/` | ✅       | Generated stubs reviewed like interface code |
| `.venv/`, `__pycache__/`      | 🚫      | Local environment only                       |

Sample `.gitignore`:

```gitignore
generated_fastapi/
__pycache__/
*.pyc
.venv/
.vscode/
```

## What You Get

| Today                                                            | Tomorrow                                                                |
| ---------------------------------------------------------------- | ----------------------------------------------------------------------- |
| ✅ A runnable FastAPI server scaffold verifying your OpenAPI spec | ✅ A fully automated, spec-driven pipeline                               |
| Code generation via OpenAPI Generator                            | Regeneration via `datamodel-codegen` + Jinja templates                  |
| Manual test of routes                                            | Contract testing with Schemathesis                                      |
| Generated “as-is” code                                           | Clean separation between generated stubs and your own `services/` logic |

## Summary

You now have a **reproducible, spec-first FastAPI code-generation pipeline** that you can re-run at any time, validate in CI, and extend safely without hand-editing generated code.
