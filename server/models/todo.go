package models

type Todo struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Done   bool   `json:"done"`
}

var todos = []Todo{
    {ID: 1, Title: "Learn Go", Done: false},
    {ID: 2, Title: "Build a ToDo App", Done: false},
}

func GetAllTodos() []Todo {
    return todos
}

func CreateTodo(todo Todo) {
    todo.ID = len(todos) + 1
    todos = append(todos, todo)
}

func UpdateTodo(id int, updated Todo) {
    for i, t := range todos {
        if t.ID == id {
            todos[i].Title = updated.Title
            todos[i].Done = updated.Done
            break
        }
    }
}

func DeleteTodo(id int) {
    for i, t := range todos {
        if t.ID == id {
            todos = append(todos[:i], todos[i+1:]...)
            break
        }
    }
}
