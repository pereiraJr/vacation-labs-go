package data

import (
	"database/sql"
	"log"
	"todo/entity"
	_ "github.com/lib/pq"
	"os"
)

func getConection() (*sql.DB, error){
	return sql.Open("postgres",os.Getenv("DATABASE_URL"))
}

func InitDatabase() {

	log.Print("db file created")

	dataBase, err := getConection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dataBase.Close()

	createTodoTable(dataBase)
}

func createTodoTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS todo (
		id SERIAL PRIMARY KEY,		
		description VARCHAR (50),
		finished BOOLEAN	
	  );`

	statemant, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	statemant.Exec()
	log.Print("todo table created")
}

func InsertTodo(todo entity.Todo) {
	db, _ := getConection()
	defer db.Close()

	query := `INSERT INTO todo (description, finished) VALUES ($1,$2);`
	statemant, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = statemant.Exec(todo.Description, todo.Finished)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func FindTodo() []entity.Todo {
	db, _ := getConection()
	defer db.Close()

	var todoList []entity.Todo
	row, err := db.Query("SELECT * FROM todo")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer row.Close()

	for row.Next() {
		var todo entity.Todo

		row.Scan(
			&todo.Id,
			&todo.Description,
			&todo.Finished,
		)
		todoList = append(todoList, todo)
	}
	return todoList
}
