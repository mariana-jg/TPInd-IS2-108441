package services

import "fmt"

type CourseNotFoundError struct {
	ID int
}

func (e *CourseNotFoundError) Error() string {
	return fmt.Sprintf("Course with ID %d not found", e.ID)
}

type CourseIncompleteError struct {
	Message string
}

func (e *CourseIncompleteError) Error() string {
	return fmt.Sprintf("Course is incomplete, you need a title and a description: %s", e.Message)
}

type RepositoryError struct {
	Message string
}

func (e *RepositoryError) Error() string {
	return e.Message
}
