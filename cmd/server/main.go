package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sarathsp06/pasteit/internal/store"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/docgen"

	"github.com/go-chi/chi"

	"github.com/sarathsp06/pasteit/internal/pasteit"
	"github.com/sarathsp06/pasteit/rpc/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	storage := store.NewMemory()
	runServer(fmt.Sprintf(":%s", port), storage)
}

func runServer(addr string, storage store.Store) {
	srv := pasteit.NewServer(1, 1024*1024, storage)
	twirpHandler := server.NewPasterServer(srv, nil)
	routes := chi.NewRouter()
	routes.Use(
		middleware.Recoverer,
		middleware.RequestID,
	)
	routes.Mount("/", twirpHandler)
	server := http.Server{
		Addr:    addr,
		Handler: routes,
	}
	log.Printf("Starting server at %q\n\nRoutes:\n", addr)
	docgen.PrintRoutes(routes)
	if err := server.ListenAndServe(); err != nil {
		log.Println("error occurred", err)
	}
}
