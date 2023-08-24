package searchinput

import (
	"log"
	"regexp"
)

const WEATHER_API = "https://api.weatherapi.com/v1"

type InputType int

const (
	Zip InputType = iota
	AirportCode
)

type WeatherInput struct {
	Input string
	Type  InputType
}

func ProcessInput(wi WeatherInput) {

}

func ValidateZipCode(input string) bool {
	validateZip, err := regexp.MatchString("/^\\d{5}$", input)

	if !validateZip || err != nil {
		log.Fatal("zip code is not valid")
		return false
	}

	return true
}

func ValidateAirportCode(input string) bool {
	validateAircode, err := regexp.MatchString("/^[A-Za-z]{3}$", input)

	if !validateAircode || err != nil {
		log.Fatal("airport code is not valid")
		return false
	}

	return true
}

func ValidateCity(input string) bool {

}
