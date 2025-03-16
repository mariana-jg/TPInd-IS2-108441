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

func TestGetCoursesController(t *testing.T) {
	mockService := services.NewMockCourseService()
	handler := controller.NewCourseHandler(mockService)
	router := setupTestRouter(handler)

	req, _ := http.NewRequest("GET", "/courses", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCourseHandlerOK(t *testing.T) {
	mockService := services.NewMockCourseService()
	handler := controller.NewCourseHandler(mockService)
	router := setupTestRouter(handler)

	req, _ := http.NewRequest("GET", "/courses/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCourseHandlerNotFound(t *testing.T) {
	mockService := services.NewMockCourseService()
	handler := controller.NewCourseHandler(mockService)
	router := setupTestRouter(handler)

	req, _ := http.NewRequest("GET", "/courses/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

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

func TestDeleteCourseHandlerOK(t *testing.T) {
	mockService := services.NewMockCourseService()
	handler := controller.NewCourseHandler(mockService)
	router := setupTestRouter(handler)

	req, _ := http.NewRequest("DELETE", "/courses/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteCourseHandlerNotFound(t *testing.T) {
	mockService := services.NewMockCourseService()
	handler := controller.NewCourseHandler(mockService)
	router := setupTestRouter(handler)

	req, _ := http.NewRequest("DELETE", "/courses/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
