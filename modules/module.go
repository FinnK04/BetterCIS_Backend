package modules

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Module defines the interface that all modules must implement.
type Module interface {
	// Name returns the module name, used as route prefix (e.g., "users" → /api/users)
	Name() string
	// RegisterRoutes registers all routes for this module on the given router
	RegisterRoutes(r chi.Router)
}

// Manager handles module registration and routing.
type Manager struct {
	modules []Module
	router  *chi.Mux
}

// NewManager creates a new module manager with a configured chi router.
func NewManager() *Manager {
	r := chi.NewRouter()

	// Standard middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	return &Manager{
		modules: make([]Module, 0),
		router:  r,
	}
}

// Register adds a module to the manager.
func (m *Manager) Register(mod Module) {
	m.modules = append(m.modules, mod)
}

// Router initializes all module routes and returns the configured router.
func (m *Manager) Router() *chi.Mux {
	m.router.Route("/api", func(r chi.Router) {
		for _, mod := range m.modules {
			r.Route("/"+mod.Name(), mod.RegisterRoutes)
		}
	})

	return m.router
}
