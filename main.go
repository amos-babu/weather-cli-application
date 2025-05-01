package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

// apiKey { f6cd6c6df35c0eeca40620360f689146 }

func main() {
	const apiKey = "f6cd6c6df35c0eeca40620360f689146"
	res, err := http.Get("https://api.openweathermap.org/data/2.5/forecast?q=Nairobi&appid=" + apiKey + "&units=metric")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	name, country, temperature, condition, description :=
		weather.City.Name,
		weather.City.Country,
		weather.List[0].Main.TempC,
		weather.List[0].Weather[0].Main,
		weather.List[0].Weather[0].Description

	fmt.Printf("%s, %s: temp: %.0fC -> %s, %s",
		name,
		country,
		temperature,
		condition,
		description,
	)
}
