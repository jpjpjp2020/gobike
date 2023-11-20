package main

import (
	"math/rand"
	"sort"
	"time"
)

type Stack struct {
	FileName string    `json:"file_name"`
	PlayAt   time.Time `json:"play_at"`
}

func generateSessionStack(data DashboardData, startTime time.Time) []Stack {

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

	var fTaskDelays []int
	if data.Mode == "Busy" {

		fTaskDelays = busyModeTimingStack(data.StartDelay, data.SessionDuration)
		shuEx := stackTasks(data.EpExercises)

		for len(shuEx) < len(fTaskDelays) {
			shuEx = append(shuEx, stackTasks(data.EpExercises)...)
		}

		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		rnd.Shuffle(len(shuEx), func(i, j int) { shuEx[i], shuEx[j] = shuEx[j], shuEx[i] })

		for i, exercise := range shuEx {
			if i < len(fTaskDelays) {
				sessionStack = append(sessionStack, Stack{
					FileName: audioActions[exercise],
					PlayAt:   startTime.Add(time.Minute * time.Duration(fTaskDelays[i])),
				})
			}
		}

	} else if data.Mode == "Surprise" {

		fTaskDelays = surpriseModeTimingStack(data.StartDelay, data.SessionDuration)
		shuEx2 := stackTasks(data.EpExercises)

		for i, exercise := range shuEx2 {
			if i < len(fTaskDelays) {
				sessionStack = append(sessionStack, Stack{
					FileName: audioActions[exercise],
					PlayAt:   startTime.Add(time.Minute * time.Duration(fTaskDelays[i])),
				})
			}
		}

	}

	return sessionStack

}

// call explictly what you need
// This should work with emergency tasks
func stackTasks(exercises []string) []string {

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

func busyModeTimingStack(startDelay, sessionDuration int) []int {
	leg1, leg2, leg3 := divideIntoLegs(sessionDuration)
	taskDelays := make([]int, 0, 9)

	taskDelays = append(taskDelays, getRandomFromLeg(leg1, 3)...)
	taskDelays = append(taskDelays, getRandomFromLeg(leg2, 3)...)
	taskDelays = append(taskDelays, getRandomFromLeg(leg3, 3)...)

	// Take start delay into account per play
	for i := range taskDelays {
		taskDelays[i] += startDelay
	}

	sort.Ints(taskDelays)
	return taskDelays
}

func surpriseModeTimingStack(startDelay, sessionDuration int) []int {
	leg1, leg2, leg3 := divideIntoLegs(sessionDuration)

	task1Delay := startDelay + getRandomFromLeg(leg1, 1)[0]
	task2Delay := startDelay + getRandomFromLeg(leg2, 1)[0]
	task3Delay := startDelay + getRandomFromLeg(leg3, 1)[0]

	return []int{task1Delay, task2Delay, task3Delay}
}

func divideIntoLegs(sessionDuration int) ([]int, []int, []int) {
	var leg1, leg2, leg3 []int
	legSize := sessionDuration / 3

	for i := 1; i < sessionDuration; i++ {
		if i <= legSize {
			leg1 = append(leg1, i)
		} else if i <= 2*legSize {
			leg2 = append(leg2, i)
		} else {
			leg3 = append(leg3, i)
		}
	}

	return leg1, leg2, leg3
}

func getRandomFromLeg(leg []int, count int) []int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	randIndices := rnd.Perm(len(leg))

	var randomMinutes []int
	for i := 0; i < count; i++ {
		randomMinutes = append(randomMinutes, leg[randIndices[i]])
	}

	return randomMinutes
}
