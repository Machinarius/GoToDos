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
		todos = []*models.TodoItem{}
	} else {
		for _, todoRecord := range todoRecords {
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
