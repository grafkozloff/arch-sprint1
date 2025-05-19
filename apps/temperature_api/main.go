package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
	"strings"
)

type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
    sensorId := ""
	path := strings.Split(r.URL.Path, "/")
	if len(path) > 3 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	if len(path) == 3 {
		sensorId = path[2]
	}
	location := r.URL.Query().Get("location")

	// If no location is provided, use a default based on sensor ID
	if location == "" {
		switch sensorId {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	// If no sensor ID is provided, generate one based on location
	if sensorId == "" {
		switch location {
		case "Living Room":
			sensorId = "1"
		case "Bedroom":
			sensorId = "2"
		case "Kitchen":
			sensorId = "3"
		default:
			sensorId = "0"
		}
	}

	response := TemperatureResponse{
		Value:       rand.Float64()*100 - 50, // случайное число от -50 до +50
		Unit:        "Celsius",
		Timestamp:   time.Now(),
		Location:    location,
		Status:      "OK",
		SensorID:    sensorId,
		SensorType:  "Thermometer",
		Description: "Temperature sensor",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func main() {
    temperatureAPIHOST := getEnv("TEMPERATURE_API_HOST", "/temperature")
    temperatureAPIPORT:= getEnv("TEMPERATURE_API_PORT", "8081")
	http.HandleFunc(temperatureAPIHOST, temperatureHandler)
	http.HandleFunc(temperatureAPIHOST + "/", temperatureHandler)
	fmt.Println("Server is running on port " + temperatureAPIPORT)
	http.ListenAndServe(":" + temperatureAPIPORT, nil)
}