package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run filename.go <latitude> <longitude>")
		return
	}

	latitude := os.Args[1]
	longitude := os.Args[2]

	url := fmt.Sprintf("http://localhost:8080/weather?lat=%s&lon=%s", latitude, longitude)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Server responded with status:", response.Status)
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read server response:", err)
		return
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(data, &weatherResponse)
	if err != nil {
		fmt.Println("Failed to parse server response:", err)
		return
	}

	fmt.Printf("Temperature in %s is %.2f degrees Celsius\n", weatherResponse.Name, weatherResponse.Main.Temp-273.15)
}
