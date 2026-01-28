package users

import "github.com/go-chi/chi/v5"

// Module implements the modules.Module interface for user management.
type Module struct{}

// New creates a new users module.
func New() *Module {
	return &Module{}
}

// Name returns the module name for routing.
func (m *Module) Name() string {
	return "users"
}

// RegisterRoutes registers all user routes.
func (m *Module) RegisterRoutes(r chi.Router) {
	registerRoutes(r)
}
