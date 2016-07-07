package sim

import (
	"math/rand"
)

func reproduce(updates map[string]*Update, organism *Organism) {
	attributes := organism.Attributes
	offspringProbability := attributes.Reproductivity / 800
	if rand.Float64() < offspringProbability {
		numberOffspring := int(attributes.Reproductivity / 30)
		for i := 0; i < numberOffspring; i++ {
			offspring := NewOrganism(organism.Attributes)
			offspring.State.Position = organism.State.Position

			updates[offspring.ID] = &Update{
				ID:         offspring.ID,
				State:      offspring.State,
				Attributes: offspring.Attributes,
			}
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
