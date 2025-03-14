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
		errorResponse := models.RFCError{
			Type:     "about:blank",
			Title:    "Invalid ID",
			Status:   http.StatusBadRequest,
			Detail:   "The provided course ID is not a valid number.",
			Instance: c.Request.URL.Path,
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	error = h.CourseService.DeleteCourse(id)
	if error != nil {
		if _, ok := error.(*services.CourseNotFoundError); ok {
			errorResponse := models.RFCError{
				Type:     "about:blank",
				Title:    "Not Found",
				Status:   http.StatusNotFound,
				Detail:   "The course with the specified ID was not found.",
				Instance: c.Request.URL.Path,
			}
			c.JSON(http.StatusNotFound, errorResponse)
			return
		}

		errorResponse := models.RFCError{
			Type:     "about:blank",
			Title:    "Internal server error",
			Status:   http.StatusInternalServerError,
			Detail:   error.Error(),
			Instance: c.Request.URL.Path,
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	c.JSON(http.StatusNoContent, nil)

}

func (h *CourseHandler) GetCourseHandler(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))

	if error != nil {
		errorResponse := models.RFCError{
			Type:     "about:blank",
			Title:    "Bad request error",
			Status:   http.StatusBadRequest,
			Detail:   "Invalid ID",
			Instance: c.Request.URL.Path,
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	course, error := h.CourseService.GetCourse(id)
	if error != nil {
		if _, ok := error.(*services.CourseNotFoundError); ok {
			errorResponse := models.RFCError{
				Type:     "about:blank",
				Title:    "Not Found",
				Status:   http.StatusNotFound,
				Detail:   "The course with the specified ID was not found.",
				Instance: c.Request.URL.Path,
			}
			c.JSON(http.StatusNotFound, errorResponse)
			return
		}

		errorResponse := models.RFCError{
			Type:     "about:blank",
			Title:    "Internal server error",
			Status:   http.StatusInternalServerError,
			Detail:   error.Error(),
			Instance: c.Request.URL.Path,
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
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

	if error := c.ShouldBindJSON(&course); error != nil {
		errorResponse := models.RFCError{
			Type:     "about:blank",
			Title:    "Bad request error",
			Status:   http.StatusBadRequest,
			Detail:   "Error on JSON structure",
			Instance: c.Request.URL.Path,
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	createdCourse, error := h.CourseService.CreateCourse(course)
	if error != nil {
		errorResponse := models.RFCError{
			Type:     "about:blank",
			Title:    "Internal server error",
			Status:   http.StatusInternalServerError,
			Detail:   error.Error(),
			Instance: c.Request.URL.Path,
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": createdCourse,
	})
}
