package main

import "time"

type Stack struct {
	FileName string    `json:"file_name"`
	PlayAt   time.Time `json:"play_at"`
}

func generate_sessions_stack(data DashboardData, startTime time.Time) []Stack {

	// custom slice of filenames and playtimes
	var sessionStack []Stack

	// mp3 file names - can map and match in frontend
	var audioActions = map[string]string{
		"start":        "start.mp3",
		"end":          "end.mp3",
		"brake":        "brake.mp3",
		"swerve_right": "swerve_right.mp3",
		"swerve_left":  "swerve_left.mp3",
	}

	for _, exercise := range data.EpExercises {
		sessionStack = append(sessionStack, Stack{
			FileName: audioActions[exercise],
			PlayAt:   timeLogic, // define in funcs
		})
	}

	return sessionStack

}

/*

NEED:

timeLogic

*/
