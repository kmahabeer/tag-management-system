# generate_services.py
import yaml
import os

spec = yaml.safe_load(open("docs/api/openapi.yaml"))

os.makedirs("app/services", exist_ok=True)

template = """from sqlalchemy.orm import Session
from app.db.models import {model_name}
from app.schemas import {model_name}Create, {model_name}Update

class {class_name}Service:
    def __init__(self, db: Session):
        self.db = db

    def create(self, data: {model_name}Create):
        obj = {model_name}(**data.dict())
        self.db.add(obj)
        self.db.commit()
        self.db.refresh(obj)
        return obj

    def get(self, id):
        return self.db.query({model_name}).filter({model_name}.id == id).first()

    def list(self):
        return self.db.query({model_name}).all()

    def update(self, id, data: {model_name}Update):
        obj = self.get(id)
        if obj:
            for k, v in data.dict(exclude_unset=True).items():
                setattr(obj, k, v)
            self.db.commit()
            self.db.refresh(obj)
        return obj

    def delete(self, id):
        obj = self.get(id)
        if obj:
            self.db.delete(obj)
            self.db.commit()
"""

for path, methods in spec["paths"].items():
    tag = list(methods.values())[0]["tags"][0]
    model_name = tag.capitalize()
    class_name = f"{model_name}Service"
    with open(f"app/services/{tag}_service.py", "w") as f:
        f.write(template.format(model_name=model_name, class_name=class_name))
