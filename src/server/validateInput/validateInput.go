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
		return inputCorrect
	}

	inputCorrect = ValidateCity(input)

	return inputCorrect
}

func ValidateZipCode(input string) bool {
	validateZip, err := regexp.MatchString("^\\d{5}(?:[-\\s]\\d{4})?$", input)

	if !validateZip || err != nil {
		log.Fatal("zip code is not valid")
		return false
	}

	return true
}

func ValidateAirportCode(input string) bool {
	validateAircode, err := regexp.MatchString("^[A-Za-z]{3}$", input)

	if !validateAircode || err != nil {
		log.Fatal("airport code is not valid")
		return false
	}

	return true
}

func ValidateCity(input string) bool {
	//validateCity, err := regexp.MatchString("^[A-Za-z]+$", input)
	//validateSameNameCityDifferentState, nexterr := regexp.MatchString("^[A-Za-z]+,\\s*[A-Za-z]{2}$", input)
	validateCityUpdated, _ := regexp.MatchString("^[A-Za-z]+,?\\s*[A-Za-z]{2}$", input)

	/*
		if strings.ContainsAny(",", input) {
			if !validateSameNameCityDifferentState || nexterr != nil {
				log.Fatal("city isn't in proper format")
				return false
			}
		} else {
			if !validateCity || err != nil {
				log.Fatal("city is either empty or wrong format (A-Z) only")
				return false
			}
		}
	*/

	if !validateCityUpdated {
		log.Fatal("city isnt correct")
		return false
	}

	return true
}
