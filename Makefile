SPEC=openapi.yaml

.PHONY: validate schemas routes format run

validate:
\topenapi-spec-validator $(SPEC)

schemas:
\tdatamodel-codegen --input $(SPEC) --input-file-type openapi --output app/schemas/models.py --target-python-version 3.11 --reuse-model --strict-nullable

routes:
\tpython tools/generate_routes.py $(SPEC)

format:
\tpython -m black app tools && python -m isort app tools

run:
\tuvicorn app.main:app --reload

all: validate schemas routes format