package sim

import (
	"fmt"
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
	if organism.State.Energy > 0.5 &&
		len(perception.Organisms) == 0 &&
		len(perception.Positions) == 0 {
		// attempt to reproduce
		update.State.Type = "reproducing"
		fmt.Println("reproducing")
	} else {

		if (len(perception.DistancePairs) > 0) {

			closestPair := perception.DistancePairs[0]

			for _, pair := range perception.DistancePairs {
				if organism.ID == pair.Organism.ID {
					continue
				}
				if pair.Distance < closestPair.Distance {
					closestPair = pair
				}
			}

			if closestPair != nil {
				// move away from
				dir := organism.State.Position.Sub(closestPair.Organism.State.Position).Normalize()
				update.State.Position = organism.State.Position.Add(dir.Mul(float32(organism.Attributes.Speed)))
			}
		}
	}
}
