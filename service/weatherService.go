package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/deepak1728/models"
	"github.com/joho/godotenv"
)

type WeatherService interface {
	GetResults(city string) models.Results
}

type weatherService struct {
	results models.Results
}

func New() WeatherService {
	return &weatherService{}
}

func (w *weatherService) GetResults(city string) models.Results {

	key := loadApiKey()
	if key == "" {
		log.Print("missing key")
	}

	params := url.Values{}
	params.Add("access_key", key)
	params.Add("query", city)

	baseUrl := "http://api.weatherstack.com/current"
	url := fmt.Sprintf("%s?%s", baseUrl, params.Encode())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error().Err(err).Msg("Request failed")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("request failed")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("unable to read body")
	}

	if err := json.Unmarshal(body, &w.results); err != nil {
		log.Error().Err(err).Msg("unmarshalling failed")
	}

	return w.results
}

func loadApiKey() string {
	var apikey string

	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msgf("Error loading .env file: %v", err)
	}

	apikey = os.Getenv("API_KEY")
	if apikey == "" {
		log.Error().Err(err).Msg("API_KEY is not set in the environment")
	}

	return apikey

}
