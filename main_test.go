package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestWelcomeHandler(t *testing.T) {
	mockResponse := `{"message":"Welcome to the API"}`
	r := SetUpRouter()
	r.GET("/", WelcomeHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestNewCourse(t *testing.T) {
	r := SetUpRouter()
	r.POST("/courses", NewCourseHandler)
	course := Course{
		ID:          "999",
		Title:       "New Course",
		Description: "New Description",
	}
	jsonValue, _ := json.Marshal(course)
	req, _ := http.NewRequest("POST", "/courses", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetCoursesHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/courses", GetCoursesHandler)
	req, _ := http.NewRequest("GET", "/courses", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var companies []Course
	json.Unmarshal(w.Body.Bytes(), &companies)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, companies)
}

func TestGetCourseHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/courses/:id", GetCourseHandler)
	req, _ := http.NewRequest("GET", "/courses/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var course Course
	json.Unmarshal(w.Body.Bytes(), &course)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, course)
}

func TestDeleteCourseHandler(t *testing.T) {
	r := SetUpRouter()
	r.DELETE("/courses/:id", DeleteCourseHandler)
	req, _ := http.NewRequest("DELETE", "/courses/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
