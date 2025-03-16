package controller

/*import (
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
)*/

/*func TestGetCoursesHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := repositories.NewCourseRepository()
	service := services.NewCoursesService(*repo)
	handler := handlers.NewCourseHandler(*service)

	router := gin.Default()
	router.GET("/courses", handler.GetCoursesHandler)

	req, _ := http.NewRequest("GET", "/courses", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
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

	var response map[string]models.Course
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	createdCourse, exists := response["data"]
	assert.True(t, exists, "Expected response to contain 'data' key")

	assert.Equal(t, "Software Engineering II", createdCourse.Title)
	assert.Equal(t, "Learn how to create your own API REST", createdCourse.Description)
}

func TestGetCourseHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := repositories.NewCourseRepository()
	service := services.NewCoursesService(*repo)
	handler := handlers.NewCourseHandler(*service)

	router := gin.Default()
	router.GET("/courses/:id", handler.GetCourseHandler)

	req, _ := http.NewRequest("GET", "/courses/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteCourseHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := repositories.NewCourseRepository()
	service := services.NewCoursesService(*repo)
	handler := handlers.NewCourseHandler(*service)

	router := gin.Default()
	router.DELETE("/courses/:id", handler.DeleteCourseHandler)

	req, _ := http.NewRequest("DELETE", "/courses/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
*/
