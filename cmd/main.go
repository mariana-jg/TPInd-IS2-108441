package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"apirest-is2/internal/handlers"
	"apirest-is2/internal/repositories"
	"apirest-is2/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func WelcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the API",
	})
}

func main() {

	error := godotenv.Load()
	if error != nil {
		log.Println("Error on .env file")
	}

	env := os.Getenv("ENVIROMENT")
	if env == "" {
		env = "development"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Starting API with configuration:")
	fmt.Println("- ENVIRONMENT:", env)
	fmt.Println("- HOST:", host)
	fmt.Println("- PORT:", port)

	router := gin.Default()

	courseRepository := repositories.NewCourseRepository()
	courseService := services.NewCoursesService(*courseRepository)
	courseHandler := handlers.NewCourseHandler(*courseService)
	courses := router.Group("/courses")
	{
		courses.GET("", courseHandler.GetCoursesHandler)
		courses.GET("/:id", courseHandler.GetCourseHandler)
		courses.POST("", courseHandler.CreateCourseHandler)
		courses.DELETE("/:id", courseHandler.DeleteCourseHandler)
	}
	router.Run(host + ":" + port)
}
