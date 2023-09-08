package formatinput

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const WEATHER_API = "http//api.weatherapi.com/v1/forecast.json?key="

type WeatherURL struct {
	formattedURL string
}

func FormatWeatherApiCall(input string) string {
	apiCall := WeatherURL{}
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file: %s", err)
	}
	apiKey := os.Getenv("WEATHER_API_KEY")
	apiCall.formattedURL = fmt.Sprintf(WEATHER_API+"%s&q=%s", apiKey, input)

	return apiCall.formattedURL
}
