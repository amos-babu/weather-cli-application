# 🌦️ Weather CLI Tool

A simple command-line application built with Go to fetch and display weather forecasts using the OpenWeatherMap API.

## 📦 Features

- Get the current weather for any city.
- Displays today's and tomorrow’s forecast.
- Shows temperature, weather condition, description, and wind speed.
- Lightweight and fast.

## 🛠 Requirements

- Go (for building from source)
- OpenWeatherMap API key

## 📥 Download (Precompiled Binary)

- Download the precompiled weather.exe and place it in a folder like C:\weather-cli.
  Run it from the terminal:

```bash
cd C:\weather-cli
weather Nairobi
```

## 📦 Download

1. Go to the [Releases](https://github.com/amos-babu/weather-cli-application/releases) page.
2. Download the latest `weather.exe` file.
3. Run it from a terminal:

   ```bash
   ./weather.exe Nairobi
   ```

### Optional Add Path

- To use weather globally from any terminal:

1.  Move weather.exe to a permanent folder, e.g., C:\weather-cli
2.  Add C:\weather-cli to your Windows Environment Variables > PATH

### 📸 Sample Output

```bash
Nairobi, KE: temp: 26°C -> Clouds, broken clouds

Mon May 13 3:00PM, temp: 25°C, Clouds, scattered clouds
Mon May 13 6:00PM, temp: 22°C, Clouds, broken clouds
Tue May 14 12:00PM, temp: 27°C, Clear, clear sky
...

```

### 🧪 Development

```bash
go run main.go Nairobi
```

## 🚀 Installation

### 1. Clone the Repository

```bash
git clone https://github.com/amos-babu/weather-cli-application
cd weather-cli
```

### 2. Set .env

```bash
API_KEY=your_openweathermap_api_key
```

### 3. Build the executable

```bash
go build -o weather.exe
```

### 3. Run the App with the city

```bash
./weather.exe Nairobi
```

### 🧾 License

MIT License

Built with ❤️ using Go ( Name Suggestions are welcomed )
