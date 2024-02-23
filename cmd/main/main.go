package main

import (
	"log"
	"net/http"

	"github.com/Woutjeee/todo_api/internal"
	"github.com/go-chi/chi/v5"
)

func main() {
	const port = "8080"
	//const fileRootPath = "."

	apiCfg := internal.ApiConfig{
		Hits: 0,
	}

	chiRouter := chi.NewRouter()

	// Setup middleware
	chiRouter.Use(apiCfg.RegisterHits)
	chiRouter.Use(apiCfg.Log)

	// Setup routes
	chiRouter.Get("/metrics", apiCfg.MetricsHandler)

	// Setup api routes
	chiRouter.HandleFunc("/api/ping", internal.Ping)
	chiRouter.Get("/api/todos", internal.GetTodos)
	chiRouter.Post("/api/todos", internal.PostTodo)

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
