package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	latitude := r.URL.Query().Get("lat")
	longitude := r.URL.Query().Get("lon")

	// Fetch weather data from OpenWeatherMap API
	response, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid={ENTER YOUR API KEY}", latitude, longitude))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Unmarshal the weather response from OpenWeatherMap API
	var weatherResponse WeatherResponse
	err = json.Unmarshal(data, &weatherResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal the weather response to JSON
	responseJSON, err := json.Marshal(weatherResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type header to JSON
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON response
	w.Write(responseJSON)
}

func main() {
	http.HandleFunc("/weather", getWeather)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

