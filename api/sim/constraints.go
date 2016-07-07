package sim

import (
	"math"
)

// ApplyConstraints updates attributes by one iteration.
func ApplyConstraints(update *Update, organism *Organism) {
	updateEnergy(update, organism)
	updateState(update, organism)
}

func updateEnergy(update *Update, organism *Organism) {
	state := update.State
	sizeFactor := state.Size * 0.1
	update.State.Energy = state.Energy - (0.01 + sizeFactor)
}

func updateMaturity(update *Update, organism *Organism) {
	update.State.Maturity += organism.Attributes.GrowthRate;
	update.State.Maturity  = math.Min(1.0, update.State.Maturity)
}

func updateState(update *Update, organism *Organism) {
	state := update.State
	if state.Energy <= 0 {
		// update state is dead
		update.State.Type = "dead"
	}
}
