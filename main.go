package main

import (
	"fmt"
	"time"
)

// JSON incoming mock
type DashboardData struct {
	SessionDuration int      `json:"session_duration"`
	StartDelay      int      `json:"start_delay"`
	Mode            string   `json:"mode"`
	EpExercises     []string `json:"ep_exercises"`
}

// testing main for manual utility func testing with mock JSON data:

func main() {
	// Mock input
	mockData := DashboardData{
		SessionDuration: 30,
		StartDelay:      5,
		Mode:            "Busy",
		EpExercises:     []string{"brake", "swerve_left", "swerve_right"},
	}

	// time now as constant to generate timestamps
	startTime := time.Now()
	sessionStack := generateSessionStack(mockData, startTime)

	// Manual output test
	fmt.Println("Generated Schedule:", sessionStack)

}

/*

type DashboardData struct {
    SessionDuration int      `json:"session_duration" validate:"required,min=1,max=60"`
    StartDelay      int      `json:"start_delay" validate:"required,min=0,max=15"`
    Mode            string   `json:"mode" validate:"required,oneof=Busy Surprise"`
    EpExercises     []string `json:"ep_exercises" validate:"required,dive,oneof=brake swerve_left swerve_right"`
}

// Validate function
func (d DashboardData) Validate() error {
    // validation library or custom logic
    return nil // or an error if validation fails
}


*/

// data traffic example

// func main() {
// 	http.HandleFunc("/setup-session", handleSessionSetup) // Setting up the route

// 	fmt.Println("Server is running on http://localhost:8080")
// 	http.ListenAndServe(":8080", nil) // Starting the server
// }

// // receive and unmarshal JSON data

// func handleSessionSetup(w http.ResponseWriter, r *http.Request) {
// 	var data DashboardData

// 	// Decode the JSON body into the struct
// 	err := json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Assume the session starts now (or get the start time from the request)
// 	startTime := time.Now()

// 	// Generate the schedule based on the user's choices
// 	schedule := generateSchedule(data, startTime)

// 	// Respond with the schedule
// 	json.NewEncoder(w).Encode(schedule)
// }
