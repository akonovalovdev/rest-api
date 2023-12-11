package routes

import (
	"github.com/go-chi/chi/v5"
	"rest-api/api/handlers"
)

func RegisterRoutes(r *chi.Mux) {
	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", handlers.GetTasks)
		r.Post("/", handlers.AddTask)
		r.Put("/", handlers.UpdateTask)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetTask)
			r.Delete("/", handlers.DeleteTask)
		})
	})
}
