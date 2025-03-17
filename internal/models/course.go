// Package models has the structs that are used to represent the data in the application.
package models

// Course is the struct that represents a course in the application
type Course struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required,min=50,max=255"`
}
