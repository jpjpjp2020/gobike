package main

import (
	"math/rand"
	"time"
)

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

	// Likely need to reorder this shuEx usage and run data through all helper funcs before using in GSS - especially in the for loop
	shuEx := stack_tasks(data.EpExercises)

	for _, exercise := range shuEx {
		sessionStack = append(sessionStack, Stack{
			FileName: audioActions[exercise],
			PlayAt:   timeLogic, // define in funcs
		})
	}

	return sessionStack

}

// call explictly what you need
// This should work with emergency tasks
func stack_tasks(exercises []string) []string {

	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	if len(exercises) == 1 {

		return []string{exercises[0], exercises[0], exercises[0]}

	}

	// pesudorandom is OK for this. use math/rand
	if len(exercises) == 2 {

		exercisesCopy := make([]string, len(exercises))
		copy(exercisesCopy, exercises)

		// rand.Seed() is depreciated!
		doubleTaskIndex := rnd.Intn(len(exercisesCopy))
		doubleTask := exercisesCopy[doubleTaskIndex]

		tasks := make([]string, 3)
		tasks[0], tasks[1] = doubleTask, doubleTask
		for _, task := range exercisesCopy {
			if task != doubleTask {
				tasks[2] = task
				break
			}
		}

		rnd.Shuffle(len(tasks), func(i, j int) { tasks[i], tasks[j] = tasks[j], tasks[i] })
		return tasks

	}

	if len(exercises) == 3 {

		exercisesCopy := make([]string, len(exercises))
		copy(exercisesCopy, exercises)

		rnd.Shuffle(len(exercisesCopy), func(i, j int) { exercisesCopy[i], exercisesCopy[j] = exercisesCopy[j], exercisesCopy[i] })
		return exercisesCopy

	}

	// fallback
	return []string{}

}

func busy_mode_timing_stack() {

}

func surprise_mode_timing_stack() {

}

func divide_into_legs() {

}

func get_random_from_leg() {

}

/*

NEED:

timeLogic

// JSON incoming mock
type DashboardData struct {
	SessionDuration int      `json:"session_duration"`
	StartDelay      int      `json:"start_delay"`
	Mode            string   `json:"mode"`
	EpExercises     []string `json:"ep_exercises"`
}

mockData := DashboardData{
	SessionDuration: 30,
	StartDelay:      5,
	Mode:            "Emergency",
	EpExercises:     []string{"brake", "swerve_left", "swerve_right"},
}

*/
