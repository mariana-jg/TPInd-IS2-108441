package main

import (
	"fmt"
	"log"
	"os"

	"apirest-is2/internal/database"
	"apirest-is2/internal/handlers"
	"apirest-is2/internal/repositories"
	"apirest-is2/internal/services"
	"apirest-is2/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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

	fmt.Println("- ENVIRONMENT:", env)
	fmt.Println("- HOST:", host)
	fmt.Println("- PORT:", port)

	logger.InitLogger()

	database.InitDB()
	defer database.DB.Close()

	router := gin.Default()

	courseRepository, err := repositories.NewCourseRepository()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	courseService := services.NewCoursesService(*courseRepository)

	if err := repositories.RunMigrations(courseRepository.DB()); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

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
