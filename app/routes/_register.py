# Auto-generated imports for routers (ordered by x-tagGroups if present)
from fastapi import FastAPI

from app.routes.meta import router as meta_router

from app.routes.tags_core import router as tags_core_router

from app.routes.tags_aliases import router as tags_aliases_router

from app.routes.tags_aliases_system import router as tags_aliases_system_router

from app.routes.tags_relationships import router as tags_relationships_router

from app.routes.tags_relationships_system import router as tags_relationships_system_router

from app.routes.tags_compositions import router as tags_compositions_router

from app.routes.tags_compositions_system import router as tags_compositions_system_router

from app.routes.tags_ratings import router as tags_ratings_router

from app.routes.tags_relationship_ratings import router as tags_relationship_ratings_router

from app.routes.entities_core import router as entities_core_router

from app.routes.entities_tags import router as entities_tags_router

from app.routes.entities_relationships import router as entities_relationships_router

from app.routes.entities_relationships_system import router as entities_relationships_system_router

from app.routes.entities_relationship_ratings import router as entities_relationship_ratings_router

from app.routes.entities_relationship_ratings_system import router as entities_relationship_ratings_system_router

from app.routes.entities_ratings import router as entities_ratings_router

from app.routes.entities_ratings_system import router as entities_ratings_system_router

from app.routes.entities_purposes import router as entities_purposes_router

from app.routes.entities_purposes_system import router as entities_purposes_system_router

from app.routes.entities_versions import router as entities_versions_router

from app.routes.entities_versions_system import router as entities_versions_system_router

from app.routes.utilities_part_of_speech import router as utilities_part_of_speech_router

from app.routes.utilities_contexts import router as utilities_contexts_router

from app.routes.utilities_ratings_system import router as utilities_ratings_system_router

from app.routes.ui_layouts import router as ui_layouts_router

from app.routes.ui_groups import router as ui_groups_router

from app.routes.ui_fields import router as ui_fields_router


def register_routers(app: FastAPI) -> None:
    
    app.include_router(meta_router)
    
    app.include_router(tags_core_router)
    
    app.include_router(tags_aliases_router)
    
    app.include_router(tags_aliases_system_router)
    
    app.include_router(tags_relationships_router)
    
    app.include_router(tags_relationships_system_router)
    
    app.include_router(tags_compositions_router)
    
    app.include_router(tags_compositions_system_router)
    
    app.include_router(tags_ratings_router)
    
    app.include_router(tags_relationship_ratings_router)
    
    app.include_router(entities_core_router)
    
    app.include_router(entities_tags_router)
    
    app.include_router(entities_relationships_router)
    
    app.include_router(entities_relationships_system_router)
    
    app.include_router(entities_relationship_ratings_router)
    
    app.include_router(entities_relationship_ratings_system_router)
    
    app.include_router(entities_ratings_router)
    
    app.include_router(entities_ratings_system_router)
    
    app.include_router(entities_purposes_router)
    
    app.include_router(entities_purposes_system_router)
    
    app.include_router(entities_versions_router)
    
    app.include_router(entities_versions_system_router)
    
    app.include_router(utilities_part_of_speech_router)
    
    app.include_router(utilities_contexts_router)
    
    app.include_router(utilities_ratings_system_router)
    
    app.include_router(ui_layouts_router)
    
    app.include_router(ui_groups_router)
    
    app.include_router(ui_fields_router)
    