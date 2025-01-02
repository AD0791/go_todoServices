package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ad0791/todoServices/api/v1/schema"
	"github.com/ad0791/todoServices/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/tools/go/analysis/passes/defers"
)

func getCOnfig() (*config.Config, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Errorf("Config error from Todo services: %v", err)
		return nil, err
	}
	return cfg, nil
}

func FetchTodos() ([]schema.TodoResponse, error) {
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

	var todos []schema.TodoResponse
	if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
		log.Errorf("fetch error for todos decoding: %v", err)
		return nil, err
	}
	return todos, nil
}

func FetchTodoByID(id string) (schema.TodoResponse, error) {
	cfg, err := getCOnfig()
	if err != nil {
		return nil, err
	}

	fetchQueryString := fmt.Sprintf("%s/%s", cfg.API.JsonPlaceholder, id)

	resp, err := http.Get(fetchQueryString)
	if err != nil {
		log.Errorf("There is an issue with fetching client: %v", err)
		return schema.TodoResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == fiber.StatusNotFound {
		log.Info("Todo "+ string(id)+ " does not exist")
		return  schema.TodoResponse{}, errors.New("todo not found")}
	}

	var todo schema.TodoResponse
	if err:= json.NewDecoder(resp.Body).Decode(&todo); err != nil{
		log.Errorf("There is an issue with fetching the todo: %v", err)
		return schema.TodoResponse{}, err
	}

	return todo, nil

}

func CreateTodo() {}

func UpdateTodoByID() {}

func DeleteTodoByID() {}
