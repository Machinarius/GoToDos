package store

import (
	"machinarius.github.io/gotodos/internal/models"
)

func GetAll() []models.TodoItem {
	return []models.TodoItem{
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
}
