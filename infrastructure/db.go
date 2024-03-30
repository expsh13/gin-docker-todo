package infrastructure

import (
	"fmt"
	"gin-docker-todo/domain"

	"github.com/jinzhu/gorm"
)

func DbInit() {
	db, err := gorm.Open("mysql", "mysql.todo")
	if err != nil {
		fmt.Errorf("could not open database")
	}
	db.AutoMigrate(&domain.Todo{})
	defer db.Close()
}

func DbCreate(todo domain.Todo) {
	db, err := gorm.Open("mysql", "mysql.todo")
	if err != nil {
		fmt.Errorf("could not open database")
	}
	db.Create(&todo)
	defer db.Close()
}

func DbRead(id ...int) []domain.Todo {
	db, err := gorm.Open("mysql", "mysql.todo")
	defer db.Close()
	if err != nil {
		fmt.Errorf("could not open database")
	}
	var todos []domain.Todo
	db.Find(&todos)
	return todos
}

func DbUpdate(id int, text string, status domain.Status, deadline int) {
	db, err := gorm.Open("mysql", "mysql.todo")
	defer db.Close()

	if err != nil {
		fmt.Errorf("could not open database")
	}
	var todo domain.Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	todo.Deadline = deadline
	db.Save(&todo)
}

func DbDelete(id int) {
	db, err := gorm.Open("mysql", "mysql.todo")
	defer db.Close()

	if err != nil {
		fmt.Errorf("could not open database")
	}
	var todo domain.Todo
	db.First(&todo, id)
	db.Delete(&todo)
}
