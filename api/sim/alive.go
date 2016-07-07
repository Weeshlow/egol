package sim

import (
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
	} else {

		var targetOrganism *Organism
		var targetPosition mgl32.Vec3

		score := 0.0;
		if len(perception.Organisms) > 0 {
			// organisms in sight
			bestScore := 0.0;
			for _, other := range perception.Organisms {
				score += 1
				score += 1 - other.Distance
				if score > bestScore {
					targetOrganism = other.Organism
					targetPosition = other.Organism.State.Position
					bestScore = score
				}
			}

		} else if len(perception.Positions) > 0 {
			// positions in sight
			bestScore := 0.0;
			for _, other := range perception.Positions {
				score += 1
				score += 1 - other.Distance
				if score > bestScore {
					targetPosition = other.Position
					bestScore = score
				}
			}
		} else if len(perception.Directions) > 0 {
			// directions
			targetPosition = perception.Directions[0]
		}

		runAway := false
		if targetOrganism != nil {
			// Check if we should be running away
			if targetOrganism.State.Energy > organism.State.Energy {
				runAway = true;
			}
		}
		dir := organism.State.Position.Sub(targetPosition).Normalize()
		if runAway {
			// move away from
			update.State.Position = organism.State.Position.Sub(dir.Mul(float32(organism.Attributes.Speed)))
		} else {
			// Chase after
			update.State.Position = organism.State.Position.Add(dir.Mul(float32(organism.Attributes.Speed)))
		}
	}
}
