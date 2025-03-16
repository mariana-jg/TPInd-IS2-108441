package services

/*import (
	"apirest-is2/internal/repositories"
	"apirest-is2/internal/services"
	"testing"
)
*/
/*func TestGetCoursesService(t *testing.T) {
	repo := repositories.NewCourseRepository()
	service := services.NewCoursesService(*repo)
	courses := service.GetCourses()
	if len(courses) != 2 {
		t.Errorf("FAIL: Expected 2 courses, got %d", len(courses))
	}
}*/
/*
func TestGetCourseService(t *testing.T) {
	repo := repositories.NewCourseRepository()
	service := services.NewCoursesService(*repo)
	course, _ := service.GetCourse(1)
	if course.ID != 1 {
		t.Errorf("FAIL: Expected course with ID 1, got %d", course.ID)
	}
}*/

/*func TestCreateCourseService(t *testing.T) {
	repo := repositories.NewCourseRepository()
	service := services.NewCoursesService(*repo)
	course := service.GetCourses()[0]
	service.CreateCourse(course)
	courses := service.GetCourses()
	if len(courses) != 3 {
		t.Errorf("FAIL: Expected 3 courses, got %d", len(courses))
	}
}

func TestDeleteCourseService(t *testing.T) {
	repo := repositories.NewCourseRepository()
	service := services.NewCoursesService(*repo)
	service.DeleteCourse(1)
	courses := service.GetCourses()
	if len(courses) != 1 {
		t.Errorf("FAIL: Expected 1 course, got %d", len(courses))
	}
}*/
