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

// Api endpoint to get all the Todos in json format.
func GetTodos(w http.ResponseWriter, r *http.Request) {
	log.Print(len(Todos))
	byteArrContent, err := json.Marshal(Todos)

	// If we have no todos, or there was an error parsing the json.
	// TODO: Do better logging here what actually is the problem.
	if len(Todos) == 0 || err != nil {
		w.Header().Set("Content-Type", "plain/text; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("There has been an error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteArrContent)
}

// API Endpoint to create Todos via json.
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

// Endpoint used to get the form or create a new todo via the form.
func CreateTodoTemplateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := make(map[string]interface{})

	// Check if we are posting, if so the user has entered some values.
	// TODO: Error handling when none values entered.
	if r.Method == "POST" {
		var todo Todo
		r.ParseForm()

		todo.Title = r.PostForm.Get("title")
		todo.Description = r.PostForm.Get("description")

		Todos = append(Todos, todo)
		CreateLog("Todo Created", Info)
		ctx["success"] = "Success"
	}
	executeTemplate("../../templates/pages/create_todo.html", w, ctx)
}
