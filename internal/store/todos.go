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

	// Pre-allocating the slice with the right size is more efficient.
	todos := make([]*models.TodoItem, len(todoRecords))
	// Mapping is idiomatically done in a loop, rather than a map function!
	// The same applies to other higher-order functions like `filter`.
	for i := range todoRecords {
		todos[i] = unwrapRecord(&todoRecords[i])
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
