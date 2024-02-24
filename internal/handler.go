package internal

import (
	"net/http"
)

func MiddlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := make(map[string]interface{})
	ctx["Name"] = "Wout"
	executeTemplate("../../templates/index.html", w, ctx)
}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := make(map[string]interface{})
	ctx["todos"] = Todos
	ctx["heading"] = "Todos"
	executeTemplate("../../templates/pages/todo.html", w, ctx)
}
