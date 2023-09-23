package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/MD-2016/Weather-App/src/server/model"

	"github.com/MD-2016/Weather-App/src/server/formatinput"
)

type FormInput struct {
	UserInput string
}

func main() {

	// get the user input

	// validate the user input

	// format the api request

	// get return object

	// display results on city page
	weatherResponse := model.Weather{}
	tmpl := template.Must(template.ParseFiles("../../../pages/index.html"))

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			tmpl.Execute(w, nil)
			log.Fatal(errors.New("wrong request being used"))
			return
		}

		var userInput FormInput
		userInput.UserInput = template.HTMLEscapeString(r.Form.Get("searchBox"))

		formattedCall := formatinput.FormatWeatherApiCall(userInput.UserInput)

		resp, err := http.Get(formattedCall)

		if err != nil {
			log.Fatal(err)
			return
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
	http.ListenAndServe(":8080", nil)
}
