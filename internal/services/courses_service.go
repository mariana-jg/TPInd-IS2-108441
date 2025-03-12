package services

import (
	"apirest-is2/internal/models"
	"apirest-is2/internal/repositories"
)

type CoursesService struct {
	repo repositories.CourseRepository
}

func NewCoursesService(repo repositories.CourseRepository) *CoursesService {
	return &CoursesService{repo}
}

func (s *CoursesService) GetCourses() []models.Course {
	return s.repo.GetCourses()
}

func (s *CoursesService) GetCourse(id int) (models.Course, error) {
	return s.repo.GetCourse(id)
}

func (s *CoursesService) CreateCourse(course models.Course) error {
	return s.repo.CreateCourse(course)
}

func (s *CoursesService) DeleteCourse(id int) error {
	return s.repo.DeleteCourse(id)
}
