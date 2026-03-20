package api

import (
	"homestack/internal/api/handlers"
	"homestack/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter builds and returns the application Chi router.
func NewRouter(svc *service.ExampleService) *chi.Mux {
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware)
	r.Use(rateLimiter(100, 200))

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/health", handlers.Health)

		example := handlers.NewExample(svc)
		r.Get("/users", example.ListUsers)
		r.Post("/users", example.CreateUser)
		r.Delete("/users/{id}", example.DeleteUser)
	})

	return r
}
