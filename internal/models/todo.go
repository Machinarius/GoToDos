package models

import "time"

// The uppercase T denotes this is a public type.
type TodoItem struct {
	// Contrary to my first impresion ID should be wholly uppercase
	// to be idiomatic to Go.
	// Go uses uppercase for the whole abbreviation by convention.
	ID          uint       `json:"id"`
	Text        string     `json:"title"`
	CompletedAt *time.Time `json:"completedAt"`
}
