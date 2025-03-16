package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"apirest-is2/internal/controller"
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

func TestGetCoursesHandler(t *testing.T) {
	mockService := services.NewMockCourseService()
	handler := controller.NewCourseHandler(mockService)

	router := setupTestRouter(handler)

	req, _ := http.NewRequest("GET", "/courses", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
