package sim

import (
	"math"
	"math/rand"
)

const (
	baseEnergyCost = 0.2
	baseAttempts = 5
)

func reproduce(updates map[string]*Update, organism *Organism) {
	attributes := organism.Attributes
	energyCost := (baseEnergyCost * attributes.Reproductivity)
	for i := 0; i < baseAttempts; i++ {
		if rand.Float64() < attributes.Reproductivity {
			// successfully create child
			offspring := NewOrganism(organism.Attributes)
			offspring.State.Position = organism.State.Position
			updates[offspring.ID] = &Update{
				ID:         offspring.ID,
				State:      offspring.State,
				Attributes: offspring.Attributes,
			}
			updates[organism.ID].State.Energy = math.Max(0, updates[organism.ID].State.Energy-energyCost)
		}
	}
}

// ReproduceAI processes the organism for the given state
func ReproduceAI(update *Update, updates map[string]*Update, organism *Organism, perception *PerceptionResults) {
	if organism.State.Energy > 0.5 &&
		len(perception.Organisms) == 0 &&
		len(perception.Positions) == 0 {
		// keep reproducing
		reproduce(updates, organism)
	} else {
		// change state back to alive
		update.State.Type = "alive";
	}
}
