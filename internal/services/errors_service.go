package services

import "fmt"

type CourseNotFoundError struct {
	ID int
}

func (e *CourseNotFoundError) Error() string {
	return fmt.Sprintf("Course with ID %d not found", e.ID)
}
