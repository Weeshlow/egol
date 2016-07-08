package sim

import (
	"math"
)

const (
	energyDepletion = 0.001
	growth          = 0.04
)

// ApplyConstraints updates attributes by one iteration.
func ApplyConstraints(update *Update) {
	updateEnergy(update)
	updateMaturity(update)
}

func updateEnergy(update *Update) {
	state := update.State
	update.State.Energy = update.State.Energy - energyDepletion
	if state.Energy <= 0 {
		// update state is dead
		update.State.Type = "dead"
		update.State.Energy = 0.0
	}
}

func updateMaturity(update *Update) {
	update.State.Maturity += growth
	// clamp
	update.State.Maturity = math.Min(1.0, update.State.Maturity)
}
