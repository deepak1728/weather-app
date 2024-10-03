package main

import (
	"github.com/deepak1728/controller"
	"github.com/deepak1728/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	weatherService := service.New()
	weatherController := controller.NewLocation(weatherService)

	current := router.Group("/current")

	current.GET("", weatherController.CurrentLocation)

	router.Run(":8080")
}
