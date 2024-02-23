package internal

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var Todos = []Todo{
	{
		Title:       "test",
		Description: "test",
	},
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	log.Print(len(Todos))
	b, err := json.Marshal(Todos)

	if len(Todos) == 0 || err != nil {
		w.Header().Set("Content-Type", "plain/text; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("There has been an error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func PostTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}

	Todos = append(Todos, todo)
	w.WriteHeader(http.StatusCreated)
	res, _ := json.Marshal(todo)
	w.Write(res)
}

func ParseTodoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := make(map[string]interface{})

	if r.Method == "POST" {
		var todo Todo
		r.ParseForm()

		todo.Title = r.PostForm.Get("title")
		todo.Description = r.PostForm.Get("description")

		Todos = append(Todos, todo)
		ctx["success"] = "Success"
	}

	t, err := template.ParseFiles("../../templates/pages/create_todo.html")
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
