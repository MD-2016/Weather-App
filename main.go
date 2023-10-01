package main

import (
	"encoding/json"
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
	http.HandleFunc("/", start)
	styles := http.FileServer(http.Dir("./src/assets/styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", styles))
	http.HandleFunc("/search", search)
	http.ListenAndServe(":8080", nil)
}

func start(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./src/pages/index.html"))
	tmpl.Execute(w, nil)
}

func search(w http.ResponseWriter, r *http.Request) {

	weatherResponse := model.Weather{}
	tmpl := template.Must(template.ParseFiles("./src/pages/index.html"))

	tmpl.Execute(w, nil)

	r.ParseForm()
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
		return
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		log.Fatal("error with decoding json call")
		return
	}

	http.Redirect(w, r, "./src/pages/city.html", 0)
	temp := template.Must(template.ParseFiles("./src/pages/city.html"))

	temp.Execute(w, weatherResponse)
	return
}
