package sim

import (
	"math"
)

const (
	energyDepletionRate = 0.005
	growthRate          = 0.04
)

func updateEnergy(update *Update) {
	update.State.Energy -= energyDepletionRate
	if update.State.Energy <= 0 {
		// update state is dead
		update.State.Type = "dead"
		// clamp
		update.State.Energy = 0.0
	}
}

func updateMaturity(update *Update) {
	update.State.Maturity += growthRate
	// clamp
	update.State.Maturity = math.Min(1.0, update.State.Maturity)
}

// ApplyConstraints updates attributes by one iteration.
func ApplyConstraints(update *Update) {
	updateEnergy(update)
	updateMaturity(update)
}
