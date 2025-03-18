// Package services contains the business logic of the application
// it ensures data validation and enforces rules before data
// is passed to repositories for persistence.
package services

import "fmt"

// CourseNotFoundError is the error that is returned when the course is not found
type CourseNotFoundError struct {
	ID int
}

// Error returns the error message
func (e *CourseNotFoundError) Error() string {
	return fmt.Sprintf("Course with ID %d not found", e.ID)
}

// CourseIncompleteError is the error that is returned when the course is incomplete
type CourseIncompleteError struct {
	Message string
}

// Error returns the error message
func (e *CourseIncompleteError) Error() string {
	return fmt.Sprintf("Course is incomplete, you need a title and a description: %s", e.Message)
}

type CourseDescriptionError struct {
	Message string
}

func (e *CourseDescriptionError) Error() string {
	return fmt.Sprintf("Course description must be between 50 and 255 characters: %s", e.Message)
}

// RepositoryError is the error that is returned when there is an error in the repository
type RepositoryError struct {
	Message string
}

// Error returns the error message
func (e *RepositoryError) Error() string {
	return e.Message
}
