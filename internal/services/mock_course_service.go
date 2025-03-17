package services

import (
	"apirest-is2/internal/models"
)

// MockCourseService is an implementation of the CoursesServiceInterface
// that uses mock data to simulate the work of the database to test the application

type MockCourseService struct {
	Courses []models.Course
}

var _ CoursesServiceInterface = (*MockCourseService)(nil)

func NewMockCourseService() *MockCourseService {
	return &MockCourseService{
		Courses: []models.Course{
			{ID: 1, Title: "Sofware Engineer I", Description: "Learn the basis of software engineering"},
			{ID: 2, Title: "Sofware Engineer II", Description: "Learn the advanced topics of software engineering"},
		},
	}
}

func (m *MockCourseService) GetCourses() ([]models.Course, error) {
	return m.Courses, nil
}

func (m *MockCourseService) GetCourse(id int) (models.Course, error) {
	for _, course := range m.Courses {
		if course.ID == id {
			return course, nil
		}
	}
	return models.Course{}, &CourseNotFoundError{ID: id}
}

func (m *MockCourseService) CreateCourse(course models.Course) (models.Course, error) {
	course.ID = len(m.Courses) + 1
	m.Courses = append(m.Courses, course)
	return course, nil
}

func (m *MockCourseService) DeleteCourse(id int) error {
	for i, course := range m.Courses {
		if course.ID == id {
			m.Courses = append(m.Courses[:i], m.Courses[i+1:]...)
			return nil
		}
	}
	return &CourseNotFoundError{ID: id}
}
