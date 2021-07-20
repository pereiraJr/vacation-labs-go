package route

import (
	"net/http"
	"todo/handler"
)

func RegisterRoute() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Health)
	mux.HandleFunc("/todo", handler.GetTodo)
	mux.HandleFunc("/todo/post", handler.InsertTodo)
	mux.HandleFunc("/todo/set-finished", handler.SetTodoFinished)

	return mux
}
