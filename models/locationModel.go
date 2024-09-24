package models

type Results struct {
	Request  RequestData  `json:"request"`
	Location LocationData `json:"location"`
	Current  CurrentData  `json:"current"`
}

type RequestData struct {
	Type     string `json:"type"`
	Query    string `json:"query"`
	Language string `json:"language"`
	Unit     string `json:"unit"`
}

type LocationData struct {
	Name           string `json:"name"`
	Country        string `json:"country"`
	Region         string `json:"region"`
	Lat            string `json:"lat"`
	Lon            string `json:"lon"`
	TimezoneId     string `json:"timezone_id"`
	LocalTime      string `json:"localtime"`
	LocalTimeEpoch string `json:"localtime_epoch"`
	UtcOffset      string `json:"utc_offset"`
}

type CurrentData struct {
	ObservationTime     string   `json:"observation_time"`
	Temperature         int      `json:"temperature"`
	WeatherCode         int      `json:"weather_code"`
	WeatherIcons        []string `json:"weather_icons"`
	WeatherDescriptions []string `json:"weather_descriptions"`
	WindSpeed           int      `json:"wind_speed"`
	WindDegree          int      `json:"wind_degree"`
	WindDir             string   `json:"wind_dir"`
	Pressure            int      `json:"pressure"`
	Precip              int      `json:"precip"`
	Humidity            int      `json:"humidity"`
	CloudCover          int      `json:"cloudcover"`
	FeelsLike           int      `json:"feelslike"`
	UvIndex             int      `json:"uv_index"`
	Visibility          int      `json:"visibility"`
	IsDay               string   `json:"is_day"`
}
