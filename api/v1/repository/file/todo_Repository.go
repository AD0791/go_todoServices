package filerepository

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"sync"

	filemodel "github.com/ad0791/todoServices/api/v1/models/file_model"
	"github.com/ad0791/todoServices/api/v1/schema"
	"github.com/gofiber/fiber/v2/log"
)

var (
	todosFile   = "data/todos.json"
	todos       []*filemodel.Todo
	simpleTodos []*schema.TodoResponse
	m           sync.Mutex
)

func loadTodos() error {
	m.Lock()
	defer m.Unlock()

	file, err := os.Open(todosFile)
	if err != nil {
		log.Errorf("cant open file to load todos: %v", err)
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(&todos)
}

func simpleLoadTodos() error {
	m.Lock()
	defer m.Unlock()

	file, err := os.Open(todosFile)
	if err != nil {
		log.Errorf("cant open file to load todos: %v", err)
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(&simpleTodos)
}

func simpleSaveTodos() error {
	m.Lock()
	defer m.Unlock()

	file, err := os.Create(todosFile)
	if err != nil {
		log.Errorf("the data file has issue: %v", err)
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(&simpleTodos)
}

func saveTodos() error {
	m.Lock()
	defer m.Unlock()

	file, err := os.Create(todosFile)
	if err != nil {
		log.Errorf("the data file has issue: %v", err)
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(&todos)
}

func GetTodos() ([]*filemodel.Todo, error) {
	if err := loadTodos(); err != nil {
		log.Errorf("We cant load the todos: %v", err)
		return nil, err
	}
	return todos, nil
}

func GetTodoByID(id string) (*filemodel.Todo, error) {
	if err := loadTodos(); err != nil {
		log.Errorf("We cant load the todos: %v", err)
		return nil, err
	}

	for _, todo := range todos {
		if strconv.Itoa(todo.ID) == id {
			log.Infof("We get the todo with the id: %s", id)
			return todo, nil
		}
	}
	return nil, errors.New("todo not found")
}

func CreateTodo(todo *filemodel.Todo) (*filemodel.Todo, error) {
	if err := loadTodos(); err != nil {
		log.Errorf("We cant load the todos: %v", err)
		return nil, err
	}

	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	if err := saveTodos(); err != nil {
		log.Errorf("we can't save/write todo for creation: %v", err)
		return nil, err
	}
	return todo, nil
}

func UpdateTodoByID(id string, updated *schema.TodoResponse) (*schema.TodoResponse, error) {
	if err := simpleLoadTodos(); err != nil {
		log.Errorf("We cant load the todos: %v", err)
		return nil, err
	}

	for i, todo := range simpleTodos {
		if strconv.Itoa(todo.ID) == id {
			updated.ID = todo.ID
			simpleTodos[i] = updated
			if err := simpleSaveTodos(); err != nil {
				log.Errorf("we can't save/write to update todo: %v", err)
				return nil, err
			}
			return simpleTodos[i], nil
		}
	}
	return nil, errors.New("todo not found for update")
}

func DeleteTodoByID(id string) (*schema.MessageResponse, error) {
	if err := loadTodos(); err != nil {
		log.Errorf("We cant load the todos: %v", err)
		return nil, err
	}

	for i, todo := range todos {
		if strconv.Itoa(todo.ID) == id {
			todos = append(todos[:i], todos[i+1:]...)
			saveTodos()
			return &schema.MessageResponse{
				ID:      todo.ID,
				Message: "Todo has been deleted",
			}, nil
		}
	}
	return nil, errors.New("todo not found for deletion")

}
