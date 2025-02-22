package player

import (
	"zahir/data"
)

var CurrentState = data.PlayerState{
	Running:      false,
	CurrentIdx:   0,
	Queue:        []int{1, 2},
	StepDuration: 0.5,
	OnEndStop:    true,
	Step:         0,
	Updated:      0,
}

var desiredState = CurrentState

func StateHasChanged() bool {
	return desiredState.Updated != CurrentState.Updated
}

func getUpdatedPlayerState() *data.PlayerState {
	if !StateHasChanged() {
		return &CurrentState
	}
	CurrentState.Running = desiredState.Running
	CurrentState.CurrentIdx = desiredState.CurrentIdx
	CurrentState.Queue = desiredState.Queue
	CurrentState.StepDuration = desiredState.StepDuration
	CurrentState.OnEndStop = desiredState.OnEndStop
	CurrentState.Step = desiredState.Step
	CurrentState.Updated = desiredState.Updated
	return &CurrentState
}

func GetCurrentPlayerState() *data.PlayerState {
	return &CurrentState
}

func SetRunning(running bool) {
	desiredState.Running = running
	desiredState.Updated = now()
}

func RunSequence(sequenceID int) {
	desiredState.Running = true
	desiredState.Queue = []int{sequenceID}
	desiredState.CurrentIdx = 0
	desiredState.Step = 0
	desiredState.OnEndStop = true
	desiredState.Updated = now()
}

func SetStepDuration(duration float64) {
	desiredState.StepDuration = duration
	desiredState.Updated = now()
}

func SetOnEndStop(onEndStop bool) {
	desiredState.OnEndStop = onEndStop
	desiredState.Updated = now()
}
