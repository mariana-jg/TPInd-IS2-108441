package main

import (
	"fmt"
	"net/http"

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
	router.Run(":8080")
}
