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
		Mode:            "Emergency",
		EpExercises:     []string{"brake", "swerve_left", "swerve_right"},
	}

	// session start
	startTime := time.Now()

	// Call your helper function with mock data
	schedule := generate_sessions_stack(mockData, startTime)

	// Manual output test
	fmt.Println("Generated Schedule:", schedule)

}

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
