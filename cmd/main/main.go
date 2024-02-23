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
	//fsHanlder := apiCfg.RegisterHits(http.StripPrefix("/app", http.FileServer(http.Dir(fileRootPath))))

	// Setup middleware
	chiRouter.Use(apiCfg.RegisterHits)
	chiRouter.Use(apiCfg.Log)

	// Setup routes
	// chiRouter.Handle("/app", fsHanlder)
	// chiRouter.Handle("/app/*", fsHanlder)
	chiRouter.Get("/metrics", apiCfg.MetricsHandler)

	// Setup
	chiRouter.HandleFunc("/ping", internal.Ping)
	chiRouter.Get("/todos", internal.GetTodos)
	chiRouter.Post("/todos", internal.PostTodo)

	corsMux := internal.MiddlewareCors(chiRouter)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
