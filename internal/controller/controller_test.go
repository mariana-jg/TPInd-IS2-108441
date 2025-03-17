package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"apirest-is2/internal/controller"
	"apirest-is2/internal/models"
	"apirest-is2/internal/services"
)

// setupTestRouter creates a gin.Engine instance with the specified handler
func setupTestRouter(handler *controller.CourseHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	courses := router.Group("/courses")
	{
		courses.GET("", handler.GetCoursesHandler)
		courses.GET("/:id", handler.GetCourseHandler)
		courses.POST("", handler.CreateCourseHandler)
		courses.DELETE("/:id", handler.DeleteCourseHandler)
	}

	return router
}

// TestGetCoursesController verifies that the GET /courses endpoint returns a 200 status code
func TestGetCoursesController(t *testing.T) {
	mockService := services.NewMockCourseService()
	handler := controller.NewCourseHandler(mockService)
	router := setupTestRouter(handler)

	req, _ := http.NewRequest("GET", "/courses", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestGetCourseHandlerOK verifies that the GET /courses/:id endpoint returns a 200 status code
// when the course exists
func TestGetCourseHandlerOK(t *testing.T) {
	mockService := services.NewMockCourseService()
	handler := controller.NewCourseHandler(mockService)
	router := setupTestRouter(handler)

	req, _ := http.NewRequest("GET", "/courses/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestGetCourseHandlerNotFound verifies that the GET /courses/:id endpoint returns a 404 status code
// when the course does not exist
func TestGetCourseHandlerNotFound(t *testing.T) {
	mockService := services.NewMockCourseService()
	handler := controller.NewCourseHandler(mockService)
	router := setupTestRouter(handler)

	req, _ := http.NewRequest("GET", "/courses/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestCreateCourseHandlerOK verifies that the POST /courses endpoint returns a 201 status code
// when the course is created successfully
func TestCreateCourseHandlerOK(t *testing.T) {
	mockService := services.NewMockCourseService()
	handler := controller.NewCourseHandler(mockService)
	router := setupTestRouter(handler)

	newCourse := models.Course{
		Title:       "Software Engineer III",
		Description: "Learn the expert topics of software engineering and build your own API using Go",
	}
	body, _ := json.Marshal(newCourse)

	req, _ := http.NewRequest("POST", "/courses", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

// TestCreateCourseHandlerDescriptionError verifies that the POST /courses endpoint returns a 400 status code
// when the course description is too short
func TestCreateCourseHandlerDescriptionError(t *testing.T) {
	mockService := services.NewMockCourseService()
	handler := controller.NewCourseHandler(mockService)
	router := setupTestRouter(handler)

	newCourse := models.Course{
		Title:       "Software Engineer III",
		Description: "Learn the topics of software engineering",
	}
	body, _ := json.Marshal(newCourse)

	req, _ := http.NewRequest("POST", "/courses", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// TestCreateCourseHandlerTitleError verifies that the POST /courses endpoint returns a 400 status code
// when the course title is empty
func TestCreateCourseHandlerBadRequest(t *testing.T) {
	mockService := services.NewMockCourseService()
	handler := controller.NewCourseHandler(mockService)
	router := setupTestRouter(handler)

	invalidCourse := `{"title": "", "description": "Short"}`
	req, _ := http.NewRequest("POST", "/courses", bytes.NewBufferString(invalidCourse))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// TestDeleteCourseHandlerOK verifies that the DELETE /courses/:id endpoint returns a 204 status code
// when the course is deleted successfully
func TestDeleteCourseHandlerOK(t *testing.T) {
	mockService := services.NewMockCourseService()
	handler := controller.NewCourseHandler(mockService)
	router := setupTestRouter(handler)

	req, _ := http.NewRequest("DELETE", "/courses/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

// TestDeleteCourseHandlerNotFound verifies that the DELETE /courses/:id endpoint returns a 404 status code
// when the course does not exist
func TestDeleteCourseHandlerNotFound(t *testing.T) {
	mockService := services.NewMockCourseService()
	handler := controller.NewCourseHandler(mockService)
	router := setupTestRouter(handler)

	req, _ := http.NewRequest("DELETE", "/courses/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
