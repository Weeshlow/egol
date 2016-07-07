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

func clampToBounds(vec *mgl32.Vec3) {
	if vec[0] < 0 {
		vec[0] = 0
	} else if vec[0] > 1 {
		vec[0] = 1
	}
	if vec[1] < 0 {
		vec[1] = 0
	} else if vec[1] > 1 {
		vec[1] = 1
	}
	if vec[2] < 0 {
		vec[2] = 0
	} else if vec[2] > 1 {
		vec[2] = 1
	}
}

// AliveAI processes the organism for the given state
func AliveAI(update *Update, updates map[string]*Update, organism *Organism, perception *PerceptionResults) {
	if organism.State.Energy > 0.5 &&
		organism.State.Maturity > 0.5 &&
		len(perception.Organisms) == 0 &&
		len(perception.Positions) == 0 {
		// attempt to reproduce
		update.State.Type = "reproducing"
	} else {
		var targetOrganism *Organism
		var targetPosition mgl32.Vec3

		if len(perception.Organisms) > 0 {
			// organisms in sight
			bestScore := 0.0
			for _, other := range perception.Organisms {
				if other.Organism.Attributes.Family == organism.Attributes.Family {
					continue
				}
				score := 0.0
				score += 1 - other.Distance
				if score > bestScore {
					targetOrganism = other.Organism
					targetPosition = other.Organism.State.Position
					bestScore = score
				}
			}
		} else if len(perception.Positions) > 0 {
			// positions in sight
			bestScore := 0.0
			for _, other := range perception.Positions {
				score := 0.0
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
			if targetOrganism.State.Energy-organism.State.Energy > 0.01 {
				runAway = true
			}
		}
		position := organism.State.Position
		diff := position.Sub(targetPosition)
		dist := position.Len()
		if dist > 0.0 {
			dir := diff.Normalize()
			speed := float32(organism.Movement())
			velocity := dir.Mul(speed)
			if runAway {
				// move away from
				update.State.Position = organism.State.Position.Add(velocity)
			} else {
				// Chase after
				update.State.Position = organism.State.Position.Sub(velocity)
			}
			clampToBounds(&update.State.Position)
		}
	}
}
