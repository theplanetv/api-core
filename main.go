package main

import (
	"api-core/libs"
	"api-core/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	libs.GetEnv()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		service := services.JourneyTagService{}
		tags := service.GetAll()

		c.JSON(http.StatusOK, gin.H{
			"data": tags,
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
