package sim

import (
	"math"
)

const (
	energyDepletionPerMS = float64(0.05 / 1000.0)
	growthPerMS          = float64(0.2 / 1000.0)
)

// ApplyConstraints updates attributes by one iteration.
func ApplyConstraints(update *Update, organism *Organism, delta int64) {
	updateEnergy(update, organism, delta)
	updateMaturity(update, organism, delta)
}

func updateEnergy(update *Update, organism *Organism, delta int64) {
	state := update.State
	update.State.Energy = update.State.Energy - (energyDepletionPerMS * float64(delta))
	if state.Energy <= 0 {
		// update state is dead
		update.State.Type = "dead"
	}
}

func updateMaturity(update *Update, organism *Organism, delta int64) {
	update.State.Maturity += growthPerMS * float64(delta)
	// clamp
	update.State.Maturity = math.Min(1.0, update.State.Maturity)
}
