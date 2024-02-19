package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	key      string
	location string
	days     string
)

// Define the Weather structure to store the JSON response from the API.
type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct { //A list of struct
			Hour []struct {
				TimeEpoch    int64   `json:"time_epoch"`
				TempC        float64 `json:"temp_c"`
				ChangeOfRain float64 `json:"chance_of_rain"`
				Condition    struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {

	// parameters
	flag.StringVar(&key, "key", "", "This is the key for the weather API")
	flag.StringVar(&location, "location", "Bogota", "This is the location for the weather API")
	flag.StringVar(&days, "days", "1", "This is the number of days for the weather API")
	flag.Parse() //This will populate the parameters

	var weather Weather

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + key + "&q=" + location + "&days=" + days)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {

		if res.StatusCode == 403 {
			log.Fatal("Invalid API key")
			panic("Invalid API key")
		}

		log.Fatal("Unexpected status code", res.StatusCode)
		panic("Weahter API is not available")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// Assigning the JSON response to the Weather struct, using the Unmarshal method from the json package.
	err = json.Unmarshal(body, &weather)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// Declaring the variables to store the location, country, current weather and the forecast for the next 24 hours.

	location, country, current, hours := weather.Location.Name, weather.Location.Country, weather.Current, weather.Forecast.Forecastday[0].Hour

	// fmt.Printf("The weather in Bogota, Colombia is Temp and condition")
	// %s -> string
	// %.0f -> float without decimals
	fmt.Printf("The weather in %s, %s is %.0fC and %s\n",
		location,
		country,
		current.TempC,
		current.Condition.Text)

	// Loop through the hours and print the weather for each hour.
	// The _ is used to ignore the index of the array. It avoids having to declare all the variables for the returns values. It is called the blank identifier.
	fmt.Println("Hour - Temp, Rain%, Condition")
	for _, hour := range hours {

		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}

		// The message is a string that contains the time, temperature, chance of rain and condition.
		message := fmt.Sprintf("%s - %.0fC, %.0f%%, %s \n",
			date.Format("15:04"), // 15:04 is the time format
			hour.TempC,
			hour.ChangeOfRain,
			hour.Condition.Text,
		)

		if hour.ChangeOfRain < 40 {

			color.Green(message)

		} else {
			color.Red(message)
		}
	}

}

func logLevel() {
	// Command line parameters definition. The variable is setup as a parameter if you use flag package.

	path := flag.String("path", "system.log", "This is the path of the log file you want to verify")
	level := flag.String("level", "com.apple.asl", "The log level you want to filter")
	flag.Parse() //This will populate the parameters

	file, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	// This will close the file after finish the execution
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(str, *level) {
			fmt.Println(str)
		}
	}
}
