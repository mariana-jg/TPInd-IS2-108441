package repositories

import (
	"apirest-is2/internal/models"
	"errors"
)

type CourseRepository struct {
	courses []models.Course
}

var ErrCourseNotFound = errors.New("course not found")

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{courses: []models.Course{
		{
			ID:          1,
			Title:       "Course #1",
			Description: "Description 1",
		},
		{
			ID:          2,
			Title:       "Course #2",
			Description: "Description 2",
		},
	},
	}
}

func (r *CourseRepository) GetCourses() []models.Course {
	return r.courses
}

func (r *CourseRepository) GetCourse(id int) (models.Course, error) {
	for _, course := range r.courses {
		if course.ID == id {
			return course, nil
		}
	}
	return models.Course{}, ErrCourseNotFound
}

func (r *CourseRepository) CreateCourse(course models.Course) (models.Course, error) {
	course.ID = len(r.courses) + 1
	r.courses = append(r.courses, course)
	return course, nil
}

func (r *CourseRepository) DeleteCourse(id int) error {
	for i, course := range r.courses {
		if course.ID == id {
			r.courses = append(r.courses[:i], r.courses[i+1:]...)
			return nil
		}
	}
	return ErrCourseNotFound
}
