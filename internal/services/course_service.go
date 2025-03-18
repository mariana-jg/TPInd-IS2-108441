// Package services contains the business logic of the application
// it ensures data validation and enforces rules before data
// is passed to repositories for persistence.
package services

import (
	"apirest-is2/internal/models"
	"apirest-is2/internal/repositories"
)

// CoursesServiceInterface is the interface that defines the methods of the CoursesService
type CoursesServiceInterface interface {
	GetCourses() ([]models.Course, error)
	GetCourse(id int) (models.Course, error)
	CreateCourse(course models.Course) (models.Course, error)
	DeleteCourse(id int) error
}

// CoursesService is the struct that represents the service of the courses
type CoursesService struct {
	repository repositories.CoursesRepositoryInterface
}

// NewCoursesService creates a new CoursesService with the repository
func NewCoursesService(repository repositories.CoursesRepositoryInterface) *CoursesService {
	return &CoursesService{repository}
}

// GetCourses returns all the courses
func (s *CoursesService) GetCourses() ([]models.Course, error) {
	courses, err := s.repository.GetCourses()
	if err != nil {
		return nil, err
	}
	return courses, nil
}

// GetCourse returns the course with the given ID
func (s *CoursesService) GetCourse(id int) (models.Course, error) {
	course, error := s.repository.GetCourse(id)
	if error != nil {
		return models.Course{}, &CourseNotFoundError{ID: id}
	}
	return course, nil
}

// CreateCourse creates a new course
func (s *CoursesService) CreateCourse(course models.Course) (models.Course, error) {
	if course.Title == "" || course.Description == "" {
		return models.Course{}, &CourseIncompleteError{Message: "Title and Description are required"}
	}

	if len(course.Description) < 50 || len(course.Description) > 255 {
		return models.Course{}, &CourseDescriptionError{Message: "Description must be between 50 and 255 characters"}
	}

	createdCourse, err := s.repository.CreateCourse(course)
	if err != nil {
		return models.Course{}, err
	}

	return createdCourse, nil
}

// DeleteCourse deletes the course with the given ID
func (s *CoursesService) DeleteCourse(id int) error {
	err := s.repository.DeleteCourse(id)
	if err != nil {
		return &CourseNotFoundError{ID: id}
	}
	return nil
}
