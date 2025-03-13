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
		error := models.RFCError{
			Type:     "about:blank",
			Title:    "Invalid ID",
			Status:   http.StatusBadRequest,
			Detail:   "The provided course ID is not a valid number.",
			Instance: c.Request.URL.Path,
		}
		c.JSON(http.StatusBadRequest, error)
		return
	}

	if h.CourseService.DeleteCourse(id) != nil {
		// i have to create an error that identifies not found course
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(http.StatusNoContent, nil)

}

func (h *CourseHandler) GetCourseHandler(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))

	if error != nil {
		error := models.RFCError{
			Type:     "about:blank",
			Title:    "Bad request error",
			Status:   http.StatusBadRequest,
			Detail:   "Invalid ID",
			Instance: c.Request.URL.Path,
		}
		c.JSON(http.StatusBadRequest, error)
		return
	}

	course, err := h.CourseService.GetCourse(id)
	if err != nil {
		// i have to create an error that identifies not found course
		error := models.RFCError{
			Type:     "about:blank",
			Title:    "Internal server error",
			Status:   http.StatusInternalServerError,
			Detail:   err.Error(),
			Instance: c.Request.URL.Path,
		}
		c.JSON(http.StatusInternalServerError, error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": course,
	})

}

func (h *CourseHandler) GetCoursesHandler(c *gin.Context) {
	courses := h.CourseService.GetCourses()
	c.JSON(http.StatusOK, gin.H{
		"data": courses,
	})
}

func (h *CourseHandler) CreateCourseHandler(c *gin.Context) {
	var course models.Course

	if err := c.ShouldBindJSON(&course); err != nil {
		error := models.RFCError{
			Type:     "about:blank",
			Title:    "Bad request error",
			Status:   http.StatusBadRequest,
			Detail:   "Error on JSON structure",
			Instance: c.Request.URL.Path,
		}
		c.JSON(http.StatusBadRequest, error)
		return
	}

	createdCourse, err := h.CourseService.CreateCourse(course)
	if err != nil {
		error := models.RFCError{
			Type:     "about:blank",
			Title:    "Internal server error",
			Status:   http.StatusInternalServerError,
			Detail:   err.Error(),
			Instance: c.Request.URL.Path,
		}
		c.JSON(http.StatusInternalServerError, error)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": createdCourse,
	})
}
