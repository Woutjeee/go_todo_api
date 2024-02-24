package main

import (
	"log"
	"net/http"

	"github.com/Woutjeee/todo_api/internal"
	"github.com/go-chi/chi/v5"
)

func main() {
	const port = "8080"

	apiCfg := internal.ApiConfig{
		Hits: 0,
	}

	chiRouter := chi.NewRouter()

	// Setup middleware
	chiRouter.Use(apiCfg.RegisterHits)
	chiRouter.Use(apiCfg.Log)

	// Setup api routes
	apiRouter := chi.NewRouter()
	apiRouter.Get("/metrics", apiCfg.MetricsHandler)
	apiRouter.HandleFunc("/ping", internal.Ping)

	// Get todos
	apiRouter.Get("/todos", internal.GetTodos)

	// Create todo via CURL
	apiRouter.Post("/todos/create", internal.PostTodo)

	// Parse todo from form and get template.
	apiRouter.Post("/todos/parse/create", internal.CreateTodoTemplateHandler)
	apiRouter.Get("/todos/parse/create", internal.CreateTodoTemplateHandler)

	// Mount all qpiRouter requests to main router.
	chiRouter.Mount("/api", apiRouter)

	// Setup tempalte routes
	chiRouter.Get("/", internal.HomeHandler)
	chiRouter.Get("/todos", internal.TodoHandler)

	corsMux := internal.MiddlewareCors(chiRouter)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
