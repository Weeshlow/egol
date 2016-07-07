package sim

import (
	"math"
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
)

// RandomPosition returns a random vec3
func RandomPosition() mgl32.Vec3 {
	return mgl32.Vec3{
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
	}
}

// AliveAI processes the organism for the given state
func AliveAI(update *Update, updates map[string]*Update, organism *Organism, perception *PerceptionResults) {
	update.State.Position = determineNewDirection(organism, perception)
	update.State.Type = "alive"
}

func determineNewDirection(currentOrganism *Organism, perception *PerceptionResults) mgl32.Vec3 {
	var closestPair = &DistancePair{Distance: math.Inf(1)}
	for _, pair := range perception.DistancePairs {
		if pair.Distance < closestPair.Distance && pair.Distance > 0 {
			closestPair = pair
		}
	}

	if nil == closestPair.Organism.State {
		return RandomPosition()
	}

	return currentOrganism.State.Position.Sub(closestPair.Organism.State.Position)
}
