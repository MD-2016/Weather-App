package searchinput

const WEATHER_API = "https://api.weatherapi.com/v1"

type InputType int

const (
	Zip InputType = iota
	UkPostCode
	CaPostCode
	AirportCode
)

type WeatherInput struct {
	Input string
	Type  InputType
}

func ProcessInput(wi WeatherInput) {
	inputLen := len(wi.Input)

	if inputLen == 5 {
		wi.Type = Zip
	} else if inputLen == 3 {

	}
}

func EvaluateZipCode(input string) bool {}

func EvaluateAirportCode(input string) bool {}
