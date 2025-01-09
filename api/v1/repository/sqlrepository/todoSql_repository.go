package sqlrepository

import (
	"github.com/ad0791/todoServices/api/v1/database"
	sqlmodel "github.com/ad0791/todoServices/api/v1/models/sql_model"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

func GetSQLTodos() ([]*sqlmodel.Todo, error) {
	var todos []*sqlmodel.Todo
	res := database.DB.Find(&todos)
	if res.Error != nil {
		log.Errorf("Failed to fetch todos: %v", res.Error)
		return nil, res.Error
	}
	log.Infof("Fetched %d todos from DB", len(todos))
	return todos, nil

}

func GetSQLTodoByID(id uint) (*sqlmodel.Todo, error) {
	var todo *sqlmodel.Todo
	res := database.DB.First(&todo, id)
	if res.Error != nil {
		log.Errorf("Probably didn't find the todo with the id %d: %v", id, res.Error)
		return nil, res.Error
	}
	return todo, nil
}

func CreateSQLTodo(todo *sqlmodel.Todo) error {
	return database.DB.Create(&todo).Error
}

func UpdateSQLTodo(todo *sqlmodel.Todo) error {
	if err := database.DB.Save(&todo).Error; err != nil {
		log.Errorf("Failed to update todo with ID %d: %v", &todo.ID, err)
		return err
	}
	log.Infof("Successfully updated todo with ID %d", &todo.ID)
	return nil
}

func DeleteAndGetSQLTodoByID(id uint) (*sqlmodel.Todo, error) {
	var todo sqlmodel.Todo

	// Soft delete the todo and ensure it exists
	res := database.DB.Where("id = ?", id).Delete(&todo)
	if res.Error != nil {
		log.Errorf("Failed to soft-delete todo with ID %d: %v", id, res.Error)
		return nil, res.Error
	}

	// Check if a record was deleted
	if res.RowsAffected == 0 {
		log.Warnf("No todo found to delete with ID %d", id)
		return nil, gorm.ErrRecordNotFound
	}

	// Retrieve the soft-deleted record
	if err := database.DB.Unscoped().First(&todo, id).Error; err != nil {
		log.Errorf("Failed to retrieve soft-deleted todo with ID %d: %v", id, err)
		return nil, err
	}

	return &todo, nil
}
