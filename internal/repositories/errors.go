// Package repositories contains the methods that interact with the database
package repositories

import "fmt"

// CourseNotFoundError is the error that is returned when the course is not found
type CourseNotFoundError struct {
	ID int
}

// Error returns the error message
func (e *CourseNotFoundError) Error() string {
	return fmt.Sprintf("Course with ID %d not found", e.ID)
}
