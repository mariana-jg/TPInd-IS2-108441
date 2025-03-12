package main

import (
	"fmt"
	"net/http"

	"apirest-is2/internal/handlers"
	"apirest-is2/internal/repositories"
	"apirest-is2/internal/services"

	"github.com/gin-gonic/gin"
)

func WelcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the API",
	})
}

func main() {
	fmt.Println("Starting...")
	router := gin.Default()

	courseRepository := repositories.NewCourseRepository()
	courseService := services.NewCoursesService(*courseRepository)
	courseHandler := handlers.NewCourseHandler(*courseService)
	router.Group("/courses")
	{
		router.GET("", courseHandler.GetCoursesHandler)
		router.GET("/:id", courseHandler.GetCourseHandler)
		router.POST("", courseHandler.CreateCourseHandler)
		router.DELETE("/:id", courseHandler.DeleteCourseHandler)
	}
	router.Run(":8080")
}
