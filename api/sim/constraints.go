package sim

import (
	"math"
)

// ApplyConstraints updates attributes by one iteration.
func ApplyConstraints(update *Update, organism *Organism) {
	updateHunger(update, organism)
	updateEnergy(update, organism)
	updateState(update, organism)
}

func updateHunger(update *Update, organism *Organism) {
	attributes := organism.Attributes
	state := organism.State
	sizeFactor := attributes.Size * 0.001

	update.State.Hunger = state.Hunger + (0.01 + sizeFactor)
}

func updateEnergy(update *Update, organism *Organism) {
	attributes := organism.Attributes
	state := organism.State

	sizeFactor := attributes.Size * 0.001

	update.State.Energy = state.Energy - 0.01 + sizeFactor

	if state.Energy < 0.7 && state.Hunger > 0 {
		consumedHunger := math.Min(0.01, state.Hunger)
		update.State.Hunger = state.Hunger - consumedHunger
		update.State.Energy = state.Hunger + (consumedHunger * 2)
	}
}

func updateState(update *Update, organism *Organism) {
	state := update.State

	if state.Energy <= 0 && state.Hunger >= 1 {
		// update state is dead
		update.State.Type = "dead"
	}

}
