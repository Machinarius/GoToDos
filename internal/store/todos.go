package store

import (
	"fmt"
	"time"

	"machinarius.github.io/gotodos/internal/models"
	"machinarius.github.io/gotodos/internal/store/database"
)

func GetAll() ([]*models.TodoItem, error) {
	var todoRecords []database.TodoRecord
	todosResult := database.DB.Find(&todoRecords)
	if todosResult.Error != nil {
		return nil, fmt.Errorf("could not load todos: %w", todosResult.Error)
	}

	var todos []*models.TodoItem
	if len(todoRecords) == 0 {
		// The loop won't be executed if the slice is empty so we need
		// to initialize it here.
		todos = []*models.TodoItem{}
	} else {
		// Mapping is idiomatically done in a loop, rather than a map function!
		// The same applies to other higher-order functions like `filter`.
		for _, todoRecord := range todoRecords {
			// append will create a new slice if the slice is nil.
			todos = append(todos, unwrapRecord(&todoRecord))
		}
	}
	return todos, nil
}

func Create(text string) ([]*models.TodoItem, error) {
	todoRecord := database.TodoRecord{
		Text: text,
	}
	result := database.DB.Create(&todoRecord)
	if result.Error != nil {
		return nil, fmt.Errorf("could not create todo: %w", result.Error)
	}
	return GetAll()
}

func unwrapRecord(record *database.TodoRecord) *models.TodoItem {
	var completedAt *time.Time
	if record.CompletedAt.Valid {
		completedAt = &record.CompletedAt.Time
	}
	return &models.TodoItem{
		ID:          record.ID,
		Text:        record.Text,
		CompletedAt: completedAt,
	}
}
