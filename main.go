package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON incoming mock
type DashboardData struct {
	SessionDuration int      `json:"session_duration"`
	StartDelay      int      `json:"start_delay"`
	Mode            string   `json:"mode"`
	EpExercises     []string `json:"ep_exercises"`
}

func main() {
	http.HandleFunc("/setup-session", handleSessionSetup) // Setting up the route

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil) // Starting the server
}

// receive and unmarshal JSON data

func handleSessionSetup(w http.ResponseWriter, r *http.Request) {
	var data DashboardData

	// Decode the JSON body into the struct
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		// Handle the error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// SET UP additional validation or processing...
	// ...

	// Respond to the client
	fmt.Fprintf(w, "Session setup received: %+v", data)
}
