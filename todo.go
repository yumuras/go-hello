package main

import (
	"encoding/json"
	"os"
	"time"
)

const dataFile = "todos.json"

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

func load() ([]Todo, error) {
	data, err := os.ReadFile(dataFile)
	if os.IsNotExist(err) {
		return []Todo{}, nil
	}
	if err != nil {
		return nil, err
	}
	var todos []Todo
	return todos, json.Unmarshal(data, &todos)
}

func save(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

func nextID(todos []Todo) int {
	max := 0
	for _, t := range todos {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1
}
