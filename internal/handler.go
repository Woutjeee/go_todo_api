package internal

import (
	"html/template"
	"log"
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
	ctx := make(map[string]string)
	ctx["Name"] = "Wout"
	t, err := template.ParseFiles("../../templates/index.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, ctx)
	if err != nil {
		log.Println("Error in template execution:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := make(map[string]interface{})
	ctx["todos"] = Todos
	ctx["heading"] = "Todos"

	log.Println(Todos)

	t, err := template.ParseFiles("../../templates/pages/todo.html")

	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, ctx)
	if err != nil {
		log.Println("Error in template execution:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
