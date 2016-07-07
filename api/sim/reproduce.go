package sim

import (
	"math/rand"
)

// Reproduce attempts to produce offspring.
func Reproduce(updates map[string]*Update, organism *Organism) {
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
