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
