package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Weather struct {
	List []struct {
		Main struct {
			TempC    float64 `json:"temp"`
			Humidity int     `json:"humidity"`
		} `json:"main"`
		Weather []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
		} `json:"weather"`
		Wind struct {
			Speed float64 `json:"speed"`
		} `json:"wind"`
		DateTime string `json:"dt_txt"`
	} `json:"list"`
	City struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"city"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("Api key not set")
	}
	res, err := http.Get("https://api.openweathermap.org/data/2.5/forecast?q=London&appid=" + apiKey + "&units=metric")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatal("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatal(err)
	}

	name, country, temperature, condition, description :=
		weather.City.Name,
		weather.City.Country,
		weather.List[0].Main.TempC,
		weather.List[0].Weather[0].Main,
		weather.List[0].Weather[0].Description

	fmt.Println()
	fmt.Printf("%s, %s: temp: %.0fC -> %s, %s\n",
		name,
		country,
		temperature,
		condition,
		description,
	)
	fmt.Println()
	// 2025-05-01 -- 13:00pm
	for _, weatherReport := range weather.List {
		fmt.Printf("%s, temp: %.0fC, %s, %s\n",
			weatherReport.DateTime,
			weatherReport.Main.TempC,
			weatherReport.Weather[0].Main,
			weatherReport.Weather[0].Description,
		)
	}

}
