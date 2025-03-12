package handlers

import (
	"apirest-is2/internal/models"
	"apirest-is2/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	CourseService services.CoursesService
}

func NewCourseHandler(service services.CoursesService) *CourseHandler {
	return &CourseHandler{CourseService: service}
}

func (h *CourseHandler) DeleteCourseHandler(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))
	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
	}

	if h.CourseService.DeleteCourse(id) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Course deleted successfully"})

}

func (h *CourseHandler) GetCourseHandler(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course not found"})
	}
	course, err := h.CourseService.GetCourse(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	c.JSON(http.StatusOK, course)

}

func (h *CourseHandler) GetCoursesHandler(c *gin.Context) {
	courses := h.CourseService.GetCourses()
	c.JSON(http.StatusOK, courses)
}

func (h *CourseHandler) CreateCourseHandler(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request error"})
		return
	}
	createdCourse, err := h.CourseService.CreateCourse(course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

	c.JSON(http.StatusCreated, createdCourse)
}
