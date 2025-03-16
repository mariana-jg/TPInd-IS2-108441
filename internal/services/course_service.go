package services

import (
	"apirest-is2/internal/models"
	"apirest-is2/internal/repositories"
	"errors"
)

type CoursesService struct {
	repository repositories.CourseRepository
}

func NewCoursesService(repository repositories.CourseRepository) *CoursesService {
	return &CoursesService{repository}
}

func (s *CoursesService) GetCourses() ([]models.Course, error) {
	//return s.repository.GetCourses()
	courses, err := s.repository.GetCourses()
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (s *CoursesService) GetCourse(id int) (models.Course, error) {
	course, error := s.repository.GetCourse(id)
	if error != nil {
		return models.Course{}, &CourseNotFoundError{ID: id}
	}
	return course, nil
}

func (s *CoursesService) CreateCourse(course models.Course) (models.Course, error) {
	if course.Title == "" || course.Description == "" {
		return models.Course{}, errors.New("title and description are required")
	}

	createdCourse, err := s.repository.CreateCourse(course)
	if err != nil {
		return models.Course{}, err
	}

	return createdCourse, nil
}

func (s *CoursesService) DeleteCourse(id int) error {
	err := s.repository.DeleteCourse(id)
	if err != nil {
		return &CourseNotFoundError{ID: id}
	}
	return nil
}
