package model

type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
}

type Current struct {
	TempC     float64   `json:"temp_c"`
	TempF     float64   `json:"temp_f"`
	Condition Condition `json:"condition"`
}

type HourlyForecast struct {
	Time      string    `json:"time"`
	TempC     float64   `json:"temp_c"`
	TempF     float64   `json:"temp_f"`
	Condition Condition `json:"condition"`
}

type ForecastDay struct {
	Hour []HourlyForecast `json:"hour"`
}

type Forecast struct {
	ForecastDay []ForecastDay `json:"forecastday"`
}

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}

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
			Icon string `json:"icon"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				Time  string  `json:"time"`
				TempC float64 `json:"temp_c"`
				TempF float64 `json:"temp_f"`

				Condition struct {
					Text string `json:"text"`
					Icon string `json:"icon"`
				} `json:"condition"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}
*/
/*
type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempF     float64 `json:"temp_f"`
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
	}
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				Time      string  `json:"time"`
				TempC     float64 `json:"temp_c"`
				TempF     float64 `json:"temp_f"`
				Condition struct {
					Text string `json:"text"`
					Icon string `json:"icon"`
				} `json:"condition"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}
*/
