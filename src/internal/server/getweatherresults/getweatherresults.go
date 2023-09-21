package getweatherresults

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/MD-2016/Weather-App/src/internal/server/formatinput"

	"github.com/MD-2016/Weather-App/src/internal/server/model"
)

/*
type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		TempF     float64 `json:"temp_f"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				TempF     float64 `json:"temp_f"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}
*/

type FormInput struct {
	UserInput string
}

func getWeather() {
	weatherResponse := model.Weather{}
	tmpl := template.Must(template.ParseFiles("../../../pages/index.html"))

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			tmpl.Execute(w, nil)
			return
		}

		var userInput FormInput
		userInput.UserInput = template.HTMLEscapeString(r.Form.Get("searchBox"))

		formattedCall := formatinput.FormatWeatherApiCall(userInput.UserInput)

		resp, err := http.Get(formattedCall)

		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode != 200 {
			log.Fatal("unable to obtain forecast from weather api")
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
			log.Fatal("error with decoding the json call")
		}

		// successful call which means we can pass into other page
		temp := template.Must(template.ParseFiles("../../../pages/city.html"))

		temp.Execute(w, weatherResponse)
		return
	})

}
