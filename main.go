package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

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
	q := "Nairobi"

	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	res, err := http.Get("https://api.openweathermap.org/data/2.5/forecast?q=" + q + "&appid=" + apiKey + "&units=metric")
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

	currentDate := time.Now().Format("2006-01-02")
	tommorrorDate := time.Now().Add(24 * time.Hour).Format("2006-01-02")

	layout := "2006-01-02 15:04:05"

	for _, forecast := range weather.List {
		parsedTime, err := time.Parse(layout, forecast.DateTime)
		if err != nil {
			fmt.Println("Parse error:", err)
			continue
		}

		forecastDate := parsedTime.Format("2006-01-02")
		if forecastDate == currentDate || forecastDate == tommorrorDate {
			fmt.Printf("%s, temp: %.0fC, %s, %s\n",
				parsedTime.Format("Mon Jan 2 3:04PM"),
				forecast.Main.TempC,
				forecast.Weather[0].Main,
				forecast.Weather[0].Description,
			)
		}
	}

}
