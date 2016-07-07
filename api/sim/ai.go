package sim

import (
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	aliveState = "alive"
	deadState = "dead"
)

// RandomPosition returns a random vec3
func RandomPosition() mgl32.Vec3 {
	return mgl32.Vec3{
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
	}
}

// Returns results from an organisms perception test
func ApplyAI(update *Update, organism *Organism, perception *PerceptionResults) {
	if organism.State.Type == deadState {
		return
	} else {
		update.State.Position = RandomPosition()
		update.State.Type = determineNextStateType(organism.ID, organism, perception.Organisms)
	}
}

func determineNextStateType(id string, organismOfInterest *Organism, organisms []*Organism) string {
	positionOfInterest := organismOfInterest.State.Position

	if organismOfInterest.State.Type == deadState {
		return deadState
	}

	for _, organism := range organisms {
		if organism.ID == id {
			continue
		}
		closeBy := positionOfInterest.ApproxEqualThreshold(organism.State.Position, 0.6)
		if closeBy && organism.State.Type != deadState {
			return deadState
		}
	}

	return aliveState
}
