package player

import (
	"fmt"
	"zahir/data"
)

func RunCycle() {
	t := now()

	for {
		state := getUpdatedPlayerState()
		currentSequenceID := state.Queue[state.CurrentIdx]
		currentSequence, _ := data.GetSequence(currentSequenceID)

		if state.Running {
			fmt.Println("Sequence", currentSequence.Name, "Step", state.Step)

			// variables, _ := data.GetSequenceVariables(currentSequenceID)
			// for _, variable := range variables {
			// 	// get value and send
			// }

			state.Step++
			updateQueue(state, currentSequence)
		}

		t = waitStep(state.StepDuration, t)
	}
}

func updateQueue(state *data.PlayerState, currentSequence *data.Sequence) {
	if len(state.Queue) == 0 {
		state.CurrentIdx = 0
		state.Step = 0
		return
	}

	if state.Step >= currentSequence.Length {
		state.CurrentIdx++
		state.Step = 0
		if state.CurrentIdx >= len(state.Queue) {
			state.CurrentIdx = 0
			if state.OnEndStop {
				state.Running = false
			}
		}
		return
	}
}
