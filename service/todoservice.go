package service

import (
	"todo/entity"
)

var results []entity.Todo
var id int

func Inserir(todo entity.Todo) {

	id++
	todo.Id = id
	results = append(results, todo)
}

func Find() []entity.Todo {
	return results
}

func SetTodoFinished(todoID int) {

	for i, todo := range results {
		if todoID == todo.Id {
			results[i].Finished = true
		}
	}
}

func UpdateTodo(todoUpdated entity.Todo) {

	for i, todo := range results {
		if todo.Id == todoUpdated.Id {
			results[i].Description = todoUpdated.Description
		}
	}
}
