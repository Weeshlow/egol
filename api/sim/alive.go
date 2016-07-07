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
	if organism.State.Energy > 0.5 &&
		len(perception.Organisms) == 0 &&
		len(perception.Positions) == 0 {
		// attempt to reproduce
		update.State.Type = "reproducing"
	} else {
		var closestPair = &DistancePair{Distance: math.Inf(1)}
		for _, pair := range perception.DistancePairs {
			if pair.Distance < closestPair.Distance && pair.Distance > 0 && currentOrganism.ID != pair.Organism.ID {
				log.Info(pair)
				closestPair = pair
			}
		}
		log.Info(closestPair.Organism.State.Position)
		log.Info(currentOrganism.State.Position)
		return currentOrganism.State.Position.Sub(closestPair.Organism.State.Position)
	}	
}
