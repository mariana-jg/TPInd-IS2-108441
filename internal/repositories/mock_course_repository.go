package repositories

import (
	"apirest-is2/internal/models"
	"errors"
)

type MockCourseRepository struct {
	Courses []models.Course
}

var _ CoursesRepositoryInterface = (*MockCourseRepository)(nil)

func NewMockCourseRepository() *MockCourseRepository {
	return &MockCourseRepository{
		Courses: []models.Course{
			{ID: 1, Title: "Sofware Engineer I", Description: "Learn the basis of software engineering"},
			{ID: 2, Title: "Sofware Engineer II", Description: "Learn the advanced topics of software engineering"},
		},
	}
}

func (m *MockCourseRepository) GetCourses() ([]models.Course, error) {
	return m.Courses, nil
}

func (m *MockCourseRepository) GetCourse(id int) (models.Course, error) {
	for _, course := range m.Courses {
		if course.ID == id {
			return course, nil
		}
	}
	return models.Course{}, errors.New("course not found")
}

func (m *MockCourseRepository) CreateCourse(course models.Course) (models.Course, error) {
	course.ID = len(m.Courses) + 1
	m.Courses = append(m.Courses, course)
	return course, nil
}

func (m *MockCourseRepository) DeleteCourse(id int) error {
	for i, course := range m.Courses {
		if course.ID == id {
			m.Courses = append(m.Courses[:i], m.Courses[i+1:]...)
			return nil
		}
	}
	return errors.New("course not found")
}
