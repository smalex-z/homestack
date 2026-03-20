package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"homestack/internal/api"
	"homestack/internal/config"
	"homestack/internal/db"
	"homestack/internal/service"
)

//go:embed frontend/dist
var frontendFS embed.FS

func main() {
	cfg := config.Load()

	// Initialize database
	database, err := db.New(cfg.DBPath)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	// Initialize service layer
	svc := service.NewExampleService(database)

	// Build router
	r := api.NewRouter(svc)

	// Serve embedded frontend as a sub-filesystem so paths resolve correctly.
	frontend, err := fs.Sub(frontendFS, "frontend/dist")
	if err != nil {
		log.Fatalf("failed to create frontend sub-filesystem: %v", err)
	}
	fileServer := http.FileServer(http.FS(frontend))

	// Any path not matched by /api routes falls through to the SPA.
	r.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fileServer.ServeHTTP(w, r)
	}))

	log.Printf("Server starting on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
