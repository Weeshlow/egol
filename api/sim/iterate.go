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
		updates[organism.ID] = &Update{
			ID: organism.ID,
			State: &State{
				Position: setRandomPositions(),
			},
		}
	}
	return updates
}

func setRandomPositions() mgl32.Vec3 {
	return mgl32.Vec3{
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
	}
}

func determineNextStateType(organismOfInterest *Organism, organisms map[string]*Organism) string {
	positionOfInterest := organismOfInterest.State.Position

	for _, organism := range organisms {
		equality := positionOfInterest.ApproxEqual(organism.State.Position)
		if equality {
			if organismOfInterest.Attributes.Energy < organism.Attributes.Energy {
				return "dead"
			}
		}
	}

	return "alive"
}
