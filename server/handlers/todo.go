package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "todo-app/models"
    "strconv"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
    todos := models.GetAllTodos()
    json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
    var todo models.Todo
    json.NewDecoder(r.Body).Decode(&todo)
    models.CreateTodo(todo)
    json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    var todo models.Todo
    json.NewDecoder(r.Body).Decode(&todo)
    models.UpdateTodo(id, todo)
    json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    models.DeleteTodo(id)
    w.WriteHeader(http.StatusNoContent)
}
