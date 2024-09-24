package controller

import (
	"github.com/deepak1728/models"
	"github.com/deepak1728/service"
	"github.com/gin-gonic/gin"
)

type LocationController interface {
	CurrentLocation(ctx *gin.Context) models.Results
}

type location struct {
	service service.WeatherService
}

func NewLocationController(service service.WeatherService) LocationController {
	return &location{service: service}
}

func (l *location) CurrentLocation(ctx *gin.Context) models.Results {
	city := ctx.Query("query")

	return l.service.GetResults(city)
}
