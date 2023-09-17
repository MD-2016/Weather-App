package model

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempF     string `json:"temp_f`
		TempC     string `json:"temp_c`
		Condition struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
	}
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				Time      string `json:"time"`
				TempC     string `json:"temp_c"`
				TempF     string `json:"temp_f"`
				Condition struct {
					Text string `json:"text"`
					Icon string `json:"icon"`
				} `json:"condition"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}
