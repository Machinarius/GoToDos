package store

import (
	"github.com/google/uuid"
	"machinarius.github.io/gotodos/internal/models"
)

var todos = []models.TodoItem{
	{
		ID:          "46abba7e-7ed7-4f16-b1d3-8cc0cc0e4f17",
		Text:        "Buy milk",
		CompletedAt: nil,
	},
	{
		ID:          "efd7a661-6e88-403d-beca-25a3beb26948",
		Text:        "Buy eggs",
		CompletedAt: nil,
	},
}

func GetAll() []models.TodoItem {
	return todos
}

func Create(text string) []models.TodoItem {
	todos = append(todos, models.TodoItem{
		ID:          uuid.NewString(),
		Text:        text,
		CompletedAt: nil,
	})
	return todos
}
