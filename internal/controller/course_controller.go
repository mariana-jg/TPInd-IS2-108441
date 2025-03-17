// Package controller has the functions that are called by the routes,
// they are the ones that interact with the services and return the response to the client.
// The functions in this package are the ones that are tested in the main_test.go file.
package controller

import (
	"apirest-is2/internal/models"
	"apirest-is2/internal/services"
	"apirest-is2/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CourseHandler is the struct that contains the services to interact with the database
// and the logger to log the requests
type CourseHandler struct {
	CourseService services.CoursesServiceInterface
	Logger        *logrus.Logger
}

// NewCourseHandler creates a new CourseHandler with the specified services
func NewCourseHandler(service services.CoursesServiceInterface) *CourseHandler {
	return &CourseHandler{
		CourseService: service,
		Logger:        logger.Logger,
	}
}

// DeleteCourseHandler is the function that deletes a course from the database
func (h *CourseHandler) DeleteCourseHandler(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))
	if error != nil {

		h.Logger.WithFields(logrus.Fields{
			"method": "DELETE",
			"status": http.StatusBadRequest,
			"path":   c.Request.URL.Path,
		}).Warn("ID is not a valid number")

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

			h.Logger.WithFields(logrus.Fields{
				"method": "DELETE",
				"status": http.StatusNotFound,
				"path":   c.Request.URL.Path,
				"id":     id,
			}).Error("Course not found")

			errorResponse := models.NewRFCError(
				http.StatusNotFound,
				"Course not found",
				"A course with the specified ID was not found.",
				c.Request.URL.Path,
			)
			c.IndentedJSON(http.StatusNotFound, errorResponse)
			return
		}

		h.Logger.WithFields(logrus.Fields{
			"method": "DELETE",
			"status": http.StatusInternalServerError,
			"error":  error.Error(),
			"path":   c.Request.URL.Path,
		}).Error("Internal Server Error")

		errorResponse := models.NewRFCError(
			http.StatusInternalServerError,
			"Internal server error",
			error.Error(),
			c.Request.URL.Path,
		)
		c.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	h.Logger.WithFields(logrus.Fields{
		"method": "DELETE",
		"status": http.StatusNoContent,
		"path":   c.Request.URL.Path,
	}).Warn("No Content response")

	c.IndentedJSON(http.StatusNoContent, nil)

}

// GetCourseHandler is the function that gets a course from the database
func (h *CourseHandler) GetCourseHandler(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))

	if error != nil {

		h.Logger.WithFields(logrus.Fields{
			"method": "GET",
			"status": http.StatusBadRequest,
			"path":   c.Request.URL.Path,
		}).Warn("ID is not a valid number")

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

			h.Logger.WithFields(logrus.Fields{
				"method": "GET",
				"status": http.StatusNotFound,
				"path":   c.Request.URL.Path,
				"id":     id,
			}).Error("Course not found")

			errorResponse := models.NewRFCError(
				http.StatusNotFound,
				"Course not found",
				"A course with the specified ID was not found.",
				c.Request.URL.Path,
			)
			c.IndentedJSON(http.StatusNotFound, errorResponse)
			return
		}

		h.Logger.WithFields(logrus.Fields{
			"method": "GET",
			"status": http.StatusInternalServerError,
			"error":  error.Error(),
			"path":   c.Request.URL.Path,
		}).Error("Internal Server Error")

		errorResponse := models.NewRFCError(
			http.StatusInternalServerError,
			"Internal server error",
			error.Error(),
			c.Request.URL.Path,
		)
		c.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	h.Logger.WithFields(logrus.Fields{
		"method": "GET",
		"status": http.StatusOK,
		"path":   c.Request.URL.Path,
	}).Info("Request successful")

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": course,
	})

}

// GetCoursesHandler is the function that gets all the courses from the database
func (h *CourseHandler) GetCoursesHandler(c *gin.Context) {
	courses, error := h.CourseService.GetCourses()

	if error != nil {

		h.Logger.WithFields(logrus.Fields{
			"method": "GET",
			"status": http.StatusInternalServerError,
			"error":  error.Error(),
			"path":   c.Request.URL.Path,
		}).Error("Internal Server Error")

		errorResponse := models.NewRFCError(
			http.StatusInternalServerError,
			"Internal server error",
			error.Error(),
			c.Request.URL.Path,
		)
		c.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	h.Logger.WithFields(logrus.Fields{
		"method": "GET",
		"status": http.StatusOK,
		"path":   c.Request.URL.Path,
	}).Info("Request successful")

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": courses,
	})
}

// CreateCourseHandler is the function that creates a course in the database
func (h *CourseHandler) CreateCourseHandler(c *gin.Context) {
	var course models.Course

	if error := c.ShouldBindJSON(&course); error != nil {

		h.Logger.WithFields(logrus.Fields{
			"method": "POST",
			"status": http.StatusBadRequest,
			"error":  error.Error(),
			"path":   c.Request.URL.Path,
		}).Warn("Bad Request received")

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

		h.Logger.WithFields(logrus.Fields{
			"method": "POST",
			"status": http.StatusInternalServerError,
			"error":  error.Error(),
			"path":   c.Request.URL.Path,
		}).Error("Internal Server Error")

		errorResponse := models.NewRFCError(
			http.StatusInternalServerError,
			"Internal server error",
			error.Error(),
			c.Request.URL.Path,
		)
		c.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	h.Logger.WithFields(logrus.Fields{
		"method": "POST",
		"status": http.StatusCreated,
		"path":   c.Request.URL.Path,
	}).Info("Course created successfully")

	c.IndentedJSON(
		http.StatusCreated,
		gin.H{"data": createdCourse})
}
