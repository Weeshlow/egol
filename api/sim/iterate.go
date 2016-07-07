package sim

import (
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
)

// Iterate applies one iteration of AI and returns the change in state as
// a map of changes.
func Iterate(organisms map[string]*Organism) map[string]*Update {
	updates := make(map[string]*Update)
	for _, organism := range organisms {
		update := &Update{
			ID:    organism.ID,
			State: &State{},
		}
		attemptReproduction(update, &updates, organism, organisms)
		ApplyConstraints(update, organism)
		if organism.State.Type == "alive" {
			update.State.Position = RandomPosition()
		}
		updates[organism.ID] = update
	}

	return updates
}

// RandomPosition returns a random vec3
func RandomPosition() mgl32.Vec3 {
	return mgl32.Vec3{
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
	}
}

func determineNextStateType(key string, organismOfInterest *Organism, organisms map[string]*Organism) string {
	positionOfInterest := organismOfInterest.State.Position

	if organismOfInterest.State.Type == "dead" {
		return "dead"
	}

	for iterKey, organism := range organisms {
		if iterKey == key {
			continue
		}
		closeBy := positionOfInterest.ApproxEqualThreshold(organism.State.Position, 0.6)
		if closeBy && organism.State.Type != "dead" {
			return "dead"
		}
	}

	return "alive"
}

// Reproductivity determines the amount and frequency of offspring
func attemptReproduction(update *Update, updates *map[string]*Update, organism *Organism, organisms map[string]*Organism) {
	attributes := organism.Attributes
	offspringProbability := attributes.Reproductivity / 400
	
	if rand.Float64() < offspringProbability {
		numberOffspring := int(attributes.Reproductivity / 30)

		for i := 0; i < numberOffspring; i++ {
			offspring := NewOrganism(organism.Attributes);
			organisms[offspring.ID] = &offspring;
			updates[offspring.ID] = &Update{
				ID: offspring.ID,
				State: offspring.State,
				Attributes: offspring.Attributes
			}
		}
	}
}