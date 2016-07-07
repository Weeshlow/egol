package sim

import (
	"math/rand"
)

// Iterate applies one iteration of AI and returns the change in state as
// a map of changes.
func Iterate(organisms map[string]*Organism) map[string]*Update {
	updates := make(map[string]*Update)
	for _, organism := range organisms {
		// create update
		update := &Update{
			ID:    organism.ID,
			State: &State{
				Energy: organism.State.Energy,
				Size: organism.State.Size,
			},
		}

		// attempt reproduction
		attemptReproduction(update, updates, organism)
		
		// apply constraints
		ApplyConstraints(update, organism)

		// apply ai
		ApplyAI(update, organism, PerceptionTest(organism, organisms))

		updates[organism.ID] = update

	}

	return updates
}

// Reproductivity determines the amount and frequency of offspring
func attemptReproduction(update *Update, updates map[string]*Update, organism *Organism) {
	attributes := organism.Attributes
	offspringProbability := attributes.Reproductivity / 800
	
	if rand.Float64() < offspringProbability {
		numberOffspring := int(attributes.Reproductivity / 30)

		update.State.Size = update.State.Size * (1 - 0.1 * float64(numberOffspring))
		update.State.Energy = update.State.Energy - 0.1 * float64(numberOffspring)

		for i := 0; i < numberOffspring; i++ {
			offspring := SpawnChild(organism);
			
			updates[offspring.ID] = &Update{
				ID: offspring.ID,
				State: offspring.State,
				Attributes: offspring.Attributes,
			}
		}
	}
}