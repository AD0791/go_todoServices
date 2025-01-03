package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/ad0791/todoServices/api/v1/schema"
	"github.com/ad0791/todoServices/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

var (
	lastID  = 200
	idMutex sync.Mutex
)

func getCOnfig() (*config.Config, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Errorf("Config error from Todo services: %v", err)
		return nil, err
	}
	return cfg, nil
}

func FetchTodos() ([]*schema.TodoResponse, error) {
	cfg, err := getCOnfig()
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(cfg.API.JsonPlaceholder)
	if err != nil {
		log.Errorf("fetch error for todos client: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var todos []*schema.TodoResponse
	if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
		log.Errorf("fetch error for todos decoding: %v", err)
		return nil, err
	}
	return todos, nil
}

func FetchTodoByID(id string) (*schema.TodoResponse, error) {
	cfg, err := getCOnfig()
	if err != nil {
		return nil, err
	}

	fetchQueryString := fmt.Sprintf("%s/%s", cfg.API.JsonPlaceholder, id)

	resp, err := http.Get(fetchQueryString)
	if err != nil {
		log.Errorf("There is an issue with fetching client: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == fiber.StatusNotFound {
		log.Infof("Todo %s does not exist", string(id))
		return nil, fmt.Errorf("todo not found")
	}

	var todo *schema.TodoResponse
	if err := json.NewDecoder(resp.Body).Decode(&todo); err != nil {
		log.Errorf("There is an issue with fetching the todo: %v", err)
		return nil, err
	}

	return todo, nil

}

func CreateTodo(req *schema.TodoRequest) (*schema.TodoResponse, error) {
	idMutex.Lock()
	lastID++
	newID := lastID
	idMutex.Unlock()

	return &schema.TodoResponse{
		ID:        newID,
		Title:     req.Title,
		Completed: req.Completed,
	}, nil
}

func UpdateTodoByID(id string, req *schema.TodoRequest) (*schema.TodoResponse, error) {
	cfg, err := getCOnfig()
	if err != nil {
		return nil, err
	}

	fetchQueryString := fmt.Sprintf("%s/%s", cfg.API.JsonPlaceholder, id)

	resp, err := http.Get(fetchQueryString)
	if err != nil {
		log.Errorf("There is an issue with fetching client: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == fiber.StatusNotFound {
		log.Infof("Todo %s does not exist", string(id))
		return nil, fmt.Errorf("todo not found")
	}

	var todo *schema.TodoResponse
	if err := json.NewDecoder(resp.Body).Decode(&todo); err != nil {
		log.Errorf("There is an issue with fetching the todo: %v", err)
		return nil, err
	}

	todo.Title = req.Title
	todo.Completed = req.Completed

	return todo, nil

}

func DeleteTodoByID(id string) (*schema.MessageResponse, error) {
	cfg, err := getCOnfig()
	if err != nil {
		return nil, err
	}

	fetchQueryString := fmt.Sprintf("%s/%s", cfg.API.JsonPlaceholder, id)

	resp, err := http.Get(fetchQueryString)
	if err != nil {
		log.Errorf("There is an issue with fetching client: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == fiber.StatusNotFound {
		log.Infof("Todo %s does not exist", string(id))
		return nil, fmt.Errorf("todo not found")
	}

	var todo *schema.TodoResponse
	if err := json.NewDecoder(resp.Body).Decode(&todo); err != nil {
		log.Errorf("There is an issue with fetching the todo: %v", err)
		return nil, err
	}

	return &schema.MessageResponse{
		ID:      todo.ID,
		Message: "Todo has been deleted",
	}, nil
}
