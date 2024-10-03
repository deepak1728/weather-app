package controller

import (
	"github.com/deepak1728/service"
	"github.com/gin-gonic/gin"
)

type LocationController interface {
	CurrentLocation(ctx *gin.Context)
}

type Location struct {
	Service service.WeatherService
}

func NewLocation(service service.WeatherService) LocationController {
	return &Location{Service: service}
}

func NewLocationController(service service.WeatherService) LocationController {
	return &Location{Service: service}
}

func (l *Location) CurrentLocation(ctx *gin.Context) {
	city := ctx.Query("query")

	results := l.Service.GetResults(city)
	ctx.JSON(200, results)
}
