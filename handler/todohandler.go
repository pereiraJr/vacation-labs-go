package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"todo/entity"
	"todo/service"
)
// Health returns that server is Up
func Health(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(responseWriter, "Server OK 9.0")
}

func GetTodo(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		jsonBody, err := json.Marshal(service.Find())
		if err != nil {
			http.Error(responseWriter, "Error parsing todo to json", http.StatusInternalServerError)
		}
		responseWriter.Write(jsonBody)
	} else {
		http.Error(responseWriter, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func InsertTodo(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(responseWriter, "Error reading request body",
				http.StatusInternalServerError)
		}

		var todo entity.Todo
		err = json.Unmarshal(body, &todo)
		if err != nil {
			http.Error(responseWriter, "Error parsing request body",
				http.StatusInternalServerError)
			panic(err)
		}

		service.Insert(todo)

		fmt.Fprint(responseWriter, "POST done")
	} else {
		http.Error(responseWriter, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func SetTodoFinished(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "PUT" {

		query := request.URL.Query()
		todoID, _ := strconv.Atoi(query.Get("todo-id"))
		fmt.Fprint(responseWriter, todoID)

		service.SetToFinished(todoID)

		fmt.Fprint(responseWriter, "POST done")
	} else {
		http.Error(responseWriter, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
