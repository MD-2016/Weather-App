package getweatherresults

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Weather struct {
	Location struct {
		Name string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
		TempF float64 `json:"temp_f"`
		Condition struct {
			Text string `json:"text"`
		}`json:"condition"`
	}`json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64 `json:"time_epoch"`
				TempC float64 `json:"temp_c"`
				TempF float64 `json:"temp_f"`
				Condition struct {
					Text string `json:"text"`
				}`json:"condition"`
			}`json:"hour"`
		}`json:"forecastday"`
	}`json:"forecast"`
}

type FormInput struct {
	UserInput string
}

func getWeather(apiCall string) error {
	tmpl := template.Must(template.ParseFiles("../../../pages/index.html"))

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			tmpl.Execute(w, nil)
			return
		}

		resp, err := http.Get(apiCall)

		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode()

	})

}
