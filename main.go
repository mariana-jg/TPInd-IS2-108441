// main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"meesage": "Welcome to the API",
	})
}

func main() {
	fmt.Println("Starting...")
	router := gin.Default() //Default en port 8080
	router.GET("/", WelcomeHandler)
	router.Run(":8080")
}
