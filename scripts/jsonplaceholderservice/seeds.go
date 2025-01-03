package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	const jsonPlaceholderURL = "https://jsonplaceholder.typicode.com/todos"
	const dataFile = "data/todos.json"

	resp, err := http.Get(jsonPlaceholderURL)
	if err != nil {
		log.Fatalf("failed to fetch todos: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	var todos []*Todo
	if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
		log.Fatalf("Failed to decode JSON response: %v", err)
	}

	file, err := os.Create(dataFile)
	if err != nil {
		log.Fatalf("Failed to create or override the data file: %v", err)
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(&todos); err != nil {
		log.Fatalf("Failed to write todos to the file: %v", err)
	}

	log.Println("Todos successfully seeded to ", dataFile)

}
