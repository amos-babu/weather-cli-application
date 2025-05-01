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
//url { https://api.openweathermap.org/data/2.5/forecast?q=Nairobi&appid=YOUR_API_KEY }
//url { http://api.openweathermap.org/data/2.5/weather?q=Nairobi&appid=" + apiKey + "&units=metric }

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

	// fmt.Println(string(body))

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	// name, country, temperature, condition :=
	// 	weather.City,
	// 	weather.City.Country,
	// 	weather.List[0].Main.TempC,
	// 	weather.List[0].Weather[0].Description

	fmt.Print(weather.List[0].Main.TempC)
	fmt.Printf("Nairobi, KE: tempC, condition")

	// fmt.Printf(
	// 	"%s, %s, %.0fC, %s\n Longitude -> %.0f, Latitide -> %.0f\n",
	// 	name,
	// 	weatherToday.Main,
	// 	temperature,
	// 	weatherToday.Description,
	// 	coordinates.Longitude, coordinates.Latitude,
	// )
}
