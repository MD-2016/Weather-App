package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/MD-2016/Weather-App/src/server/formatinput"
	"github.com/MD-2016/Weather-App/src/server/model"
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
	http.HandleFunc("/", start)
	styles := http.FileServer(http.Dir("./src/assets/styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", styles))
	badInput := http.FileServer(http.Dir("./src/pages/404.html"))
	http.Handle("/bad", http.StripPrefix("/bad", badInput))
	http.HandleFunc("/search", searchHandler)
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

	} else {
		fmt.Println(inputToParse)
		weatherRes := model.Weather{}
		formattedCall := formatinput.FormatWeatherApiCall(inputToParse)
		fmt.Println(formattedCall)
		res, err := http.Get(formattedCall)

		fmt.Println(res)

		if res.StatusCode != 200 {
			log.Fatal(err)
			return
		}
		defer res.Body.Close()

		if err := json.NewDecoder(r.Body).Decode(&weatherRes); err != nil {
			fmt.Println(err)
			log.Fatal(err)
			return
		}

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
