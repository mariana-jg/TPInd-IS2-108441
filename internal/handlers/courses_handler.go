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
		errorResponse := models.NewRFCError(
			http.StatusBadRequest,
			"Invalid ID",
			"ID is not a valid number",
			c.Request.URL.Path,
		)
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}

	error = h.CourseService.DeleteCourse(id)
	if error != nil {
		if _, ok := error.(*services.CourseNotFoundError); ok {
			errorResponse := models.NewRFCError(
				http.StatusNotFound,
				"Course not found",
				"A course with the specified ID was not found.",
				c.Request.URL.Path,
			)
			c.IndentedJSON(http.StatusNotFound, errorResponse)
			return
		}
		errorResponse := models.NewRFCError(
			http.StatusInternalServerError,
			"Internal server error",
			error.Error(),
			c.Request.URL.Path,
		)
		c.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)

}

func (h *CourseHandler) GetCourseHandler(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))

	if error != nil {
		errorResponse := models.NewRFCError(
			http.StatusBadRequest,
			"Invalid ID",
			"ID is not a valid number",
			c.Request.URL.Path,
		)
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}

	course, error := h.CourseService.GetCourse(id)
	if error != nil {
		if _, ok := error.(*services.CourseNotFoundError); ok {
			errorResponse := models.NewRFCError(
				http.StatusNotFound,
				"Course not found",
				"A course with the specified ID was not found.",
				c.Request.URL.Path,
			)
			c.IndentedJSON(http.StatusNotFound, errorResponse)
			return
		}

		errorResponse := models.NewRFCError(
			http.StatusInternalServerError,
			"Internal server error",
			error.Error(),
			c.Request.URL.Path,
		)
		c.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": course,
	})

}

func (h *CourseHandler) GetCoursesHandler(c *gin.Context) {
	courses := h.CourseService.GetCourses()
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": courses,
	})
}

func (h *CourseHandler) CreateCourseHandler(c *gin.Context) {
	var course models.Course

	if error := c.ShouldBindJSON(&course); error != nil {
		errorResponse := models.NewRFCError(
			http.StatusBadRequest,
			"Bad request error",
			"Error on JSON structure",
			c.Request.URL.Path,
		)
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}

	createdCourse, error := h.CourseService.CreateCourse(course)
	if error != nil {
		errorResponse := models.NewRFCError(
			http.StatusInternalServerError,
			"Internal server error",
			error.Error(),
			c.Request.URL.Path,
		)
		c.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"data": createdCourse,
	})
}
