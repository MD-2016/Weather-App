package searchinput

import (
	"html/template"
	"net/http"
	"strings"
	"log"
)

const WEATHER_API = "https://api.weatherapi.com/v1"

type WeatherInput struct {
	City string
	Zip int
	UKpost string
	CApost string
	AirportCode string
}


func 