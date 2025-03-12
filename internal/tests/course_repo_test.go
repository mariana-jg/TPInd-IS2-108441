package tests

import (
	"apirest-is2/internal/repositories"
	"testing"
)

// We have charged the repository with two courses, so we expect to get two courses always

func TestGetCourses(t *testing.T) {
	repo := repositories.NewCourseRepository()
	courses := repo.GetCourses()
	if len(courses) != 2 {
		t.Errorf("Expected 2 courses, got %d", len(courses))
	}
}

func TestGetCourse(t *testing.T) {
	repo := repositories.NewCourseRepository()
	course, _ := repo.GetCourse(1)
	if course.ID != 1 {
		t.Errorf("Expected course with ID 1, got %d", course.ID)
	}
}

func TestCreateCourse(t *testing.T) {
	repo := repositories.NewCourseRepository()
	course := repo.GetCourses()[0]
	repo.CreateCourse(course)
	courses := repo.GetCourses()
	if len(courses) != 3 {
		t.Errorf("Expected 3 courses, got %d", len(courses))
	}
}

func TestDeleteCourse(t *testing.T) {
	repo := repositories.NewCourseRepository()
	repo.DeleteCourse(1)
	courses := repo.GetCourses()
	if len(courses) != 1 {
		t.Errorf("Expected 1 course, got %d", len(courses))
	}
}
