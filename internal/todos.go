package internal

import (
	"encoding/json"
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
