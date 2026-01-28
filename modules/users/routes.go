package users

import "github.com/go-chi/chi/v5"

func registerRoutes(r chi.Router) {
	r.Get("/", handleGetAll)
	r.Post("/", handleCreate)
	r.Get("/{id}", handleGetByID)
	r.Put("/{id}", handleUpdate)
	r.Delete("/{id}", handleDelete)
}
