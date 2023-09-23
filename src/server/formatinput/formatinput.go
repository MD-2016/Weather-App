package formatinput

import (
	"fmt"
	"log"
	"os"

	"github.com/MD-2016/Weather-App/src/server/validateInput"
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
	if validateInput.ValidateInput(input) == false {
		log.Fatal("input is not valid for the api call")
	}
	apiKey := os.Getenv("WEATHER_API_KEY")
	apiCall.formattedURL = fmt.Sprintf(WEATHER_API+"%s&q=%s&days=10&aqi=no&alerts=no", apiKey, input)

	return apiCall.formattedURL
}
