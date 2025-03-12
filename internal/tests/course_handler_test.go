package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"apirest-is2/internal/handlers"
	"apirest-is2/internal/models"
	"apirest-is2/internal/repositories"
	"apirest-is2/internal/services"
)

func TestGetCoursesHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := repositories.NewCourseRepository()
	service := services.NewCoursesService(*repo)
	handler := handlers.NewCourseHandler(*service)

	router := gin.Default()
	router.GET("/courses", handler.GetCoursesHandler)

	req, _ := http.NewRequest("GET", "/courses", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestCreateCourseHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := repositories.NewCourseRepository()
	service := services.NewCoursesService(*repo)
	handler := handlers.NewCourseHandler(*service)

	router := gin.Default()
	router.POST("/courses", handler.CreateCourseHandler)

	course := models.Course{
		Title:       "Software Engineering II",
		Description: "Learn how to create your own API REST",
	}
	jsonData, _ := json.Marshal(course)

	req, _ := http.NewRequest("POST", "/courses", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdCourse models.Course
	json.Unmarshal(w.Body.Bytes(), &createdCourse)

	assert.Equal(t, "Software Engineering II", createdCourse.Title)
	assert.Equal(t, "Learn how to create your own API REST", createdCourse.Description)
}
