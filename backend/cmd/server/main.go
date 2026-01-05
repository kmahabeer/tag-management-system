package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/kmahabeer/tag-management-system/backend/internal/handlers"
	"github.com/kmahabeer/tag-management-system/backend/internal/middleware"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.LoggingMiddleware)
	r.Use(corsMiddleware)
	r.Use(middleware.ValidationMiddleware)
	r.Use(middleware.ErrorHandler)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/meta/health", handlers.HealthCheck)

		r.Route("/tags", func(r chi.Router) {
			r.Get("/", handlers.ListTags)
			r.Post("/", handlers.CreateTag)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetTag)
				r.Patch("/", handlers.UpdateTag)
				r.Delete("/", handlers.DeleteTag)
				r.Route("/aliases", func(r chi.Router) {
					r.Get("/", handlers.ListTagAliases)
					r.Post("/", handlers.CreateTagAlias)
					r.Patch("/", handlers.UpdateTagAliases)
					r.Delete("/", handlers.DeleteTagAliases)
					r.Route("/{alias_id}", func(r chi.Router) {
						r.Get("/", handlers.GetTagAlias)
						r.Patch("/", handlers.UpdateTagAlias)
						r.Delete("/", handlers.DeleteTagAlias)
					})
				})
				r.Route("/relationships", func(r chi.Router) {
					r.Get("/", handlers.ListTagRelationships)
					r.Post("/", handlers.CreateTagRelationship)
					r.Patch("/", handlers.UpdateTagRelationships)
					r.Delete("/", handlers.DeleteTagRelationships)
					r.Route("/{relationship_id}", func(r chi.Router) {
						r.Get("/", handlers.GetTagRelationship)
						r.Patch("/", handlers.UpdateTagRelationship)
						r.Delete("/", handlers.DeleteTagRelationship)
					})
				})
				r.Route("/compositions", func(r chi.Router) {
					r.Get("/", handlers.ListTagCompositions)
					r.Post("/", handlers.CreateTagComposition)
					r.Patch("/", handlers.UpdateTagCompositions)
					r.Delete("/", handlers.DeleteTagCompositions)
					r.Route("/{composition_id}", func(r chi.Router) {
						r.Get("/", handlers.GetTagComposition)
						r.Patch("/", handlers.UpdateTagComposition)
						r.Delete("/", handlers.DeleteTagComposition)
					})
				})
				r.Route("/ratings", func(r chi.Router) {
					r.Get("/", handlers.ListTagRatings)
					r.Patch("/", handlers.UpdateTagRatings)
					r.Delete("/", handlers.DeleteTagRatings)
					r.Route("/{rating_id}", func(r chi.Router) {
						r.Get("/", handlers.GetTagRating)
						r.Patch("/", handlers.UpdateTagRating)
						r.Delete("/", handlers.DeleteTagRating)
					})
				})
				r.Route("/relationship_ratings", func(r chi.Router) {
					r.Get("/", handlers.ListTagRelationshipRatings)
					r.Patch("/", handlers.UpdateTagRelationshipRatings)
					r.Delete("/", handlers.DeleteTagRelationshipRatings)
					r.Route("/{rating_id}", func(r chi.Router) {
						r.Get("/", handlers.GetTagRelationshipRating)
						r.Patch("/", handlers.UpdateTagRelationshipRating)
						r.Delete("/", handlers.DeleteTagRelationshipRating)
					})
				})
			})
		})
		r.Route("/tag-aliases", func(r chi.Router) {
			r.Get("/", handlers.ListTagAliasesSystem)
			r.Post("/", handlers.CreateTagAliasSystem)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetTagAliasSystem)
				r.Patch("/", handlers.UpdateTagAliasSystem)
				r.Delete("/", handlers.DeleteTagAliasSystem)
			})
		})
		r.Route("/tag-relationships", func(r chi.Router) {
			r.Get("/", handlers.ListTagRelationshipsSystem)
			r.Post("/", handlers.CreateTagRelationshipSystem)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetTagRelationshipSystem)
				r.Patch("/", handlers.UpdateTagRelationshipSystem)
				r.Delete("/", handlers.DeleteTagRelationshipSystem)
			})
		})
		r.Route("/tag-compositions", func(r chi.Router) {
			r.Get("/", handlers.ListTagCompositionsSystem)
			r.Post("/", handlers.CreateTagCompositionSystem)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetTagCompositionSystem)
				r.Patch("/", handlers.UpdateTagCompositionSystem)
				r.Delete("/", handlers.DeleteTagCompositionSystem)
			})
		})

		r.Route("/entities", func(r chi.Router) {
			r.Get("/", handlers.ListEntities)
			r.Post("/", handlers.CreateEntity)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetEntity)
				r.Patch("/", handlers.UpdateEntity)
				r.Delete("/", handlers.DeleteEntity)
				r.Route("/tags", func(r chi.Router) {
					r.Get("/", handlers.ListEntityTags)
					r.Patch("/", handlers.UpdateEntityTags)
					r.Route("/{tag_id}", func(r chi.Router) {
						r.Get("/", handlers.GetEntityTag)
						r.Delete("/", handlers.DeleteEntityTag)
					})
				})
				r.Route("/purposes", func(r chi.Router) {
					r.Get("/", handlers.ListEntityPurposes)
					r.Patch("/", handlers.UpdateEntityPurposes)
					r.Route("/{purpose_id}", func(r chi.Router) {
						r.Get("/", handlers.GetEntityPurpose)
						r.Patch("/", handlers.UpdateEntityPurpose)
						r.Delete("/", handlers.DeleteEntityPurpose)
					})
				})
				r.Route("/versions", func(r chi.Router) {
					r.Get("/", handlers.ListEntityVersions)
					r.Patch("/", handlers.UpdateEntityVersions)
					r.Route("/{version_id}", func(r chi.Router) {
						r.Get("/", handlers.GetEntityVersion)
						r.Patch("/", handlers.UpdateEntityVersion)
						r.Delete("/", handlers.DeleteEntityVersion)
					})
				})
				r.Route("/ratings", func(r chi.Router) {
					r.Get("/", handlers.ListEntityRatings)
					r.Patch("/", handlers.UpdateEntityRatings)
					r.Route("/{rating_id}", func(r chi.Router) {
						r.Get("/", handlers.GetEntityRating)
						r.Patch("/", handlers.UpdateEntityRating)
						r.Delete("/", handlers.DeleteEntityRating)
					})
				})
				r.Route("/relationships", func(r chi.Router) {
					r.Get("/", handlers.ListEntityRelationships)
					r.Patch("/", handlers.UpdateEntityRelationships)
					r.Route("/{relationship_id}", func(r chi.Router) {
						r.Get("/", handlers.GetEntityRelationship)
						r.Patch("/", handlers.UpdateEntityRelationship)
						r.Delete("/", handlers.DeleteEntityRelationship)
					})
				})
				r.Route("/relationship_ratings", func(r chi.Router) {
					r.Get("/", handlers.ListEntityRelationshipRatings)
					r.Patch("/", handlers.UpdateEntityRelationshipRatings)
					r.Route("/{rating_id}", func(r chi.Router) {
						r.Get("/", handlers.GetEntityRelationshipRating)
						r.Patch("/", handlers.UpdateEntityRelationshipRating)
						r.Delete("/", handlers.DeleteEntityRelationshipRating)
					})
				})
			})
		})
		r.Route("/entity-relationships", func(r chi.Router) {
			r.Get("/", handlers.ListEntityRelationshipsSystem)
			r.Post("/", handlers.CreateEntityRelationshipSystem)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetEntityRelationshipSystem)
				r.Patch("/", handlers.UpdateEntityRelationshipSystem)
				r.Delete("/", handlers.DeleteEntityRelationshipSystem)
			})
		})
		r.Route("/entity-relationship-ratings", func(r chi.Router) {
			r.Get("/", handlers.ListEntityRelationshipRatingsSystem)
			r.Post("/", handlers.CreateEntityRelationshipRatingSystem)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetEntityRelationshipRatingSystem)
				r.Patch("/", handlers.UpdateEntityRelationshipRatingSystem)
				r.Delete("/", handlers.DeleteEntityRelationshipRatingSystem)
			})
		})
		r.Route("/entity-ratings", func(r chi.Router) {
			r.Get("/", handlers.ListEntityRatingsSystem)
			r.Post("/", handlers.CreateEntityRatingSystem)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetEntityRatingSystem)
				r.Patch("/", handlers.UpdateEntityRatingSystem)
				r.Delete("/", handlers.DeleteEntityRatingSystem)
			})
		})
		r.Route("/entity-purposes", func(r chi.Router) {
			r.Get("/", handlers.ListEntityPurposesSystem)
			r.Post("/", handlers.CreateEntityPurposeSystem)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetEntityPurposeSystem)
				r.Patch("/", handlers.UpdateEntityPurposeSystem)
				r.Delete("/", handlers.DeleteEntityPurposeSystem)
			})
		})
		r.Route("/entity-versions", func(r chi.Router) {
			r.Get("/", handlers.ListEntityVersionsSystem)
			r.Post("/", handlers.CreateEntityVersionSystem)
		})

		r.Route("/contexts", func(r chi.Router) {
			r.Get("/", handlers.ListContexts)
			r.Post("/", handlers.CreateContext)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetContext)
				r.Patch("/", handlers.UpdateContext)
				r.Delete("/", handlers.DeleteContext)
			})
		})
		r.Route("/parts-of-speech", func(r chi.Router) {
			r.Get("/", handlers.ListPartsOfSpeech)
			r.Post("/", handlers.CreatePartOfSpeech)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetPartOfSpeech)
				r.Patch("/", handlers.UpdatePartOfSpeech)
				r.Delete("/", handlers.DeletePartOfSpeech)
			})
		})
		r.Route("/ratings", func(r chi.Router) {
			r.Get("/", handlers.ListRatings)
			r.Post("/", handlers.CreateRating)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetRating)
				r.Patch("/", handlers.UpdateRating)
				r.Delete("/", handlers.DeleteRating)
			})
		})
		r.Route("/rating-types", func(r chi.Router) {
			r.Get("/", handlers.ListRatingTypes)
			r.Post("/", handlers.CreateRatingType)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetRatingType)
				r.Patch("/", handlers.UpdateRatingType)
				r.Delete("/", handlers.DeleteRatingType)
			})
		})

		r.Route("/ui/layouts", func(r chi.Router) {
			r.Get("/", handlers.ListUILayouts)
			r.Post("/", handlers.CreateUILayout)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetUILayout)
				r.Patch("/", handlers.UpdateUILayout)
				r.Delete("/", handlers.DeleteUILayout)
			})
		})
		r.Route("/ui/groups", func(r chi.Router) {
			r.Get("/", handlers.ListUIGroups)
			r.Post("/", handlers.CreateUIGroup)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetUIGroup)
				r.Patch("/", handlers.UpdateUIGroup)
				r.Delete("/", handlers.DeleteUIGroup)
			})
		})
		r.Route("/ui/fields", func(r chi.Router) {
			r.Get("/", handlers.ListUIFields)
			r.Post("/", handlers.CreateUIField)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetUIField)
				r.Patch("/", handlers.UpdateUIField)
				r.Delete("/", handlers.DeleteUIField)
			})
		})
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Println("Server starting on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
