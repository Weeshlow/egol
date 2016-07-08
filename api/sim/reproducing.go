package sim

import (
	"math"
	"math/rand"
)

const (
	baseEnergyCost = 0.4
	baseAttempts   = 3
)

func reproduce(update *Update, updates map[string]*Update, organism *Organism) {
	attributes := organism.Attributes
	energyCost := math.Max(0.1, baseEnergyCost-attributes.Reproductivity)
	for i := 0; i < baseAttempts; i++ {
		if update.State.Energy < energyCost {
			break
		}
		if rand.Float64() < attributes.Reproductivity {
			// successfully create child
			offspring := organism.Spawn()
			// add to updates
			updates[offspring.ID] = &Update{
				ID:         offspring.ID,
				State:      offspring.State,
				Attributes: offspring.Attributes,
			}
			// reduce energy trying to reproduce
			update.State.Energy = math.Max(0, update.State.Energy-energyCost)
		}
	}
}

// ReproduceAI processes the organism for the given state.
func ReproduceAI(update *Update, updates map[string]*Update, organism *Organism, perception *PerceptionResults) {
	if organism.State.Energy > 0.5 &&
		len(perception.Threats) == 0 {
		// keep reproducing
		reproduce(update, updates, organism)
		update.State.Type = "reproducing"
	} else {
		// change state back to alive
		update.State.Type = "alive"
	}
}
