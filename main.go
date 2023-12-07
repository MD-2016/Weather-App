package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/MD-2016/Weather-App/src/server/formatinput"
	"github.com/MD-2016/Weather-App/src/server/model"
	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
)

type FormInput struct {
	UserInput string
}

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func main() {

	// rate limit for too many api calls
	message := Message{
		Status: "Request Failed too many requests",
		Body:   "The API rwached capacity, try again later",
	}

	apiError, _ := json.Marshal(message)
	//lmt := tollbooth.NewLimiter(1, nil)
	lmt := tollbooth.NewLimiter(1, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	lmt.SetMessageContentType("application/json")
	lmt.SetMessage(string(apiError))
	lmt.SetMethods([]string{"GET"})
	lmt.SetOnLimitReached(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request limit is reached")
	})

	// get the user input

	// validate the user input

	// format the api request

	// get return object

	// display results on city page
	http.HandleFunc("/", start)
	styles := http.FileServer(http.Dir("./src/assets/styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", styles))
	http.Handle("/search", tollbooth.LimitFuncHandler(lmt, searchHandler))
	http.HandleFunc("/search/{city}", searchHandler)
	//http.HandleFunc("/search/", searchCityHandler)
	http.ListenAndServe(":8080", nil)

}

func start(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./src/pages/index.html"))
	tmpl.Execute(w, nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	//tmpl := template.Must(template.ParseFiles("./src/pages/index.html"))
	inputToParse := r.URL.Query().Get("city")
	if inputToParse == "" {
		return
	} else {
		weatherRes := model.Weather{}
		formattedCall := formatinput.FormatWeatherApiCall(inputToParse)

		res, err := http.Get(formattedCall)

		if err != nil || res.StatusCode != http.StatusOK {
			log.Fatal(err)
			return
		}

		defer res.Body.Close()

		if err := json.NewDecoder(res.Body).Decode(&weatherRes); err != nil {
			log.Fatal(err)
			return
		}

		tmpl := template.Must(template.ParseFiles("./src/pages/city.html"))
		tmpl.Execute(w, weatherRes)
		return

	}

	//city := fmt.Sprintf("/search/%s", userInput.UserInput)

}

/*
func searchCityHandler(w http.ResponseWriter, r *http.Request) {
	weatherRes := model.Weather{}
	tmpl := template.Must(template.ParseFiles("./src/pages/city.html"))

	formattedCall := formatinput.FormatWeatherApiCall(r.URL.Path)

	res, err := http.Get(formattedCall)

	if res.StatusCode != 200 {
		log.Fatal(err)
		return
	}

	defer res.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&weatherRes); err != nil {
		log.Fatal(err)
		return
	}

	tmpl.Execute(w, weatherRes)
}
*/
