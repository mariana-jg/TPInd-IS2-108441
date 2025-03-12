package services

import (
	"apirest-is2/internal/models"
	"apirest-is2/internal/repositories"
)

type CoursesService struct {
	repository repositories.CourseRepository
}

func NewCoursesService(repository repositories.CourseRepository) *CoursesService {
	return &CoursesService{repository}
}

func (s *CoursesService) GetCourses() []models.Course {
	return s.repository.GetCourses()
}

func (s *CoursesService) GetCourse(id int) (models.Course, error) {
	return s.repository.GetCourse(id)
}

func (s *CoursesService) CreateCourse(course models.Course) (models.Course, error) {
	createdCourse, err := s.repository.CreateCourse(course)
	if err != nil {
		return models.Course{}, err
	}

	return createdCourse, nil
}

func (s *CoursesService) DeleteCourse(id int) error {
	return s.repository.DeleteCourse(id)
}
