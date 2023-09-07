package formatinput

import (
	"fmt"
)

const WEATHER_API = "http//api.weatherapi.com/v1"

type WeatherURL struct {
	formattedURL string
}

func FormatWeatherApiCall(string input) {
	apiCall := fmt.Sprintf(WEATHER_API + "")
}
