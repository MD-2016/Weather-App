package getweatherresults

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type WeatherResults struct {
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
