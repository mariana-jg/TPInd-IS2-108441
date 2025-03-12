package main

import (
	"fmt"
	"net/http"

	"slices"

	"github.com/gin-gonic/gin"
)

type Course struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var courses = []Course{
	{
		ID:          "1",
		Title:       "Course #1",
		Description: "Description 1",
	},
	{
		ID:          "2",
		Title:       "Course #2",
		Description: "Description 2",
	},
}

func DeleteCourseHandler(c *gin.Context) {
	id := c.Param("id")
	for i, course := range courses {
		if course.ID == id {
			courses = slices.Delete(courses, i, i+1)
			c.JSON(http.StatusOK, gin.H{"message": "Course deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
}

func GetCourseHandler(c *gin.Context) {
	id := c.Param("id")
	for _, course := range courses {
		if course.ID == id {
			c.JSON(http.StatusOK, course)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
}

func GetCoursesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, courses)
}

func NewCourseHandler(c *gin.Context) {
	var course Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	courses = append(courses, course)
	c.JSON(http.StatusCreated, course)
}

func WelcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the API",
	})
}

func main() {
	fmt.Println("Starting...")
	router := gin.Default()
	router.GET("/", WelcomeHandler)
	router.POST("/courses", NewCourseHandler)
	router.GET("/courses", GetCoursesHandler)
	router.GET("/courses/:id", GetCourseHandler)
	router.DELETE("/courses/:id", DeleteCourseHandler)
	router.Run(":8080")
}
