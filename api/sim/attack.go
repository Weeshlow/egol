package sim

import (
	"math/rand"
)

// ReproduceAI processes the organism for the given state
func AttackAI(update *Update, updates map[string]*Update, organism *Organism, perception *PerceptionResults) {
	if len(perception.Organisms) > 0 {
		// organisms in sight
		for _, other := range perception.Organisms {
			if other.Organism.Attributes.Family == organism.Attributes.Family {
				continue
			}
			if other.Distance-other.Organism.State.Size-organism.State.Size < organism.Attributes.Range {
				// in attack range, lets attack
				updates[other.Organism.ID].State.Energy -= (0.25 * rand.Float64())
				update.State.Energy += (0.25 * rand.Float64())
				return
			}
		}
	}
	// return to idle state
	update.State.Type = "alive"
}
