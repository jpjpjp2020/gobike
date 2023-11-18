package main

import "time"

type AudioCue struct {
	FileName string    `json:"file_name"`
	PlayAt   time.Time `json:"play_at"`
}

func generate_sessions_stack() {

}

func generateSchedule(data DashboardData, startTime time.Time) []AudioCue {
	var schedule []AudioCue

	// Calculate the initial start time with the delay
	currentTime := startTime.Add(time.Minute * time.Duration(data.StartDelay))

	for _, exercise := range data.EpExercises {
		// Example: Add a fixed interval between exercises
		currentTime = currentTime.Add(time.Minute * 5)

		schedule = append(schedule, AudioCue{
			FileName: exercise + ".mp3",
			PlayAt:   currentTime,
		})
	}

	return schedule
}
