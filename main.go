package main

import (
	"github.com/deepak1728/controller"
	"github.com/deepak1728/service"
	"github.com/gin-gonic/gin"
)

var (
	weatherService    service.WeatherService        = service.New()
	weatherController controller.LocationController = controller.NewLocationController(weatherService)
)

func main() {
	router := gin.Default()

	router.GET("/report", func(ctx *gin.Context) {
		ctx.JSON(200, weatherController.CurrentLocation(ctx))

	})

	router.Run(":8080")
}
