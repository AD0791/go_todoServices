package sqlrepository

import (
	"github.com/ad0791/todoServices/api/v1/database"
	sqlmodel "github.com/ad0791/todoServices/api/v1/models/sql_model"
	"github.com/gofiber/fiber/v2/log"
)

func GetSQLTodos() ([]*sqlmodel.Todo, error) {
	var todos []*sqlmodel.Todo
	res := database.DB.Find(&todos)
	if res.Error != nil {
		log.Errorf("Probably didn't find the todo %v: %v", todos, res.Error)
	}
	return todos, res.Error

}

func GetSQLTodoByID(id uint) (*sqlmodel.Todo, error) {
	var todo *sqlmodel.Todo
	res := database.DB.First(&todo, id)
	if res.Error != nil {
		log.Errorf("Probably didn't find the todo with the id %d: %v", id, res.Error)
	}
	return todo, res.Error
}

func CreateSQLTodo(todo *sqlmodel.Todo) error {
	return database.DB.Create(&todo).Error
}

func UpdateSQLTodo(todo *sqlmodel.Todo) error {
	return database.DB.Save(&todo).Error
}

func DeleteSQLTodo(id uint) error {
	return database.DB.Delete(&sqlmodel.Todo{}, id).Error
}
