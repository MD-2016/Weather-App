package validateInput

import (
	"log"
	"regexp"
	"strings"
)

/*
type InputType int

const (
	Zip InputType = iota
	AirportCode
)
*/

func ValidateInput(input string) bool {
	inputCorrect := false

	if input == "" {
		return inputCorrect
	}
	if firstChar, _ := regexp.MatchString("^[0-9]$", input[0:1]); firstChar {
		inputCorrect = ValidateZipCode(input)
		return inputCorrect
	}

	if len(input) == 3 && strings.ToUpper(input) == input {
		inputCorrect = ValidateAirportCode(input)
	}

	inputCorrect = ValidateCity(input)

	return inputCorrect
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
	validateCity, err := regexp.MatchString("^[A-Za-z ]*$", input)

	if !validateCity || err != nil {
		log.Fatal("city is not valid")
		return false
	}

	return true
}
