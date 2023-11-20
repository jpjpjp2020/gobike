package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

// JSON incoming
type DashboardData struct {
	SessionDuration int      `json:"session_duration" validate:"required,oneof=15 30 60"`
	StartDelay      int      `json:"start_delay" validate:"required,oneof=1 5 15"`
	Mode            string   `json:"mode" validate:"required,oneof=Busy Surprise"`
	EpExercises     []string `json:"ep_exercises" validate:"required,dive,oneof=brake swerve_left swerve_right"`
}

func main() {
	http.HandleFunc("/setup-session", handleSessionSetup) // Setting up the route

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil) // Starting the server
}

func handleSessionSetup(w http.ResponseWriter, r *http.Request) {
	var data DashboardData

	// Decode the JSON body into the struct
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := data.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assume the session starts now (or get the start time from the request)
	startTime := time.Now()

	// Generate the schedule based on the user's choices
	schedule := generateSessionStack(data, startTime)

	// Respond with the schedule
	json.NewEncoder(w).Encode(schedule)
}

func (d DashboardData) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
