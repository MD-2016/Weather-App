package getweatherresults

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/MD-2016/Weather-App/src/internal/server/formatinput"
)

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

type FormInput struct {
	UserInput string
}

func getWeather() {
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

		var weather Weather
		if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
			log.Fatal("error with decoding the json call")
		}

		location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

		//test code for getting the right results
		fmt.Printf(
			"%s, %s: %.0fC, \n",
			location.Name,
			location.Country,
			current.TempC,
			current.Condition.Text,
		)

		for _, hour := range hours {
			date := time.Unix(hour.TimeEpoch, 0)

			if date.Before(time.Now()) {
				continue
			}

			mess := fmt.Sprintf(
				"%s - %.0fC, %.0fF, %s\n",
				date.Format("15:04"),
				hour.TempC,
				hour.TempF,
				hour.Condition.Text,
			)

			fmt.Print(mess)
		}

	})

}
