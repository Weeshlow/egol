package sim

import (
	"math"
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/unchartedsoftware/egol/api/util"
)

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
		len(perception.Organisms) == 0 {
		// attempt to reproduce
		update.State.Type = "reproducing"
	} else {

		if len(perception.Organisms) > 0 {
			// organisms in sight
			var targetPosition *mgl32.Vec3
			shortestDist := math.MaxFloat64
			for _, other := range perception.Organisms {
				if other.Organism.Attributes.Family == organism.Attributes.Family {
					continue
				}
				if other.Distance-other.Organism.State.Size-organism.State.Size < organism.Attributes.Range {
					// in attack range, lets attack
					update.State.Type = "attack"
					return
				}
				if other.Distance < shortestDist {
					//targetOrganism = other.Organism
					targetPosition = &other.Organism.State.Position
					shortestDist = other.Distance
				}
			}

			if targetPosition != nil {
				position := organism.State.Position
				diff := targetPosition.Sub(position)
				dist := diff.Len()
				if dist > 0.0 {
					dir := diff.Normalize()
					speed := float32(organism.Movement())
					velocity := dir.Mul(speed)
					noise := util.RandomDirection().Mul(0.01)

					score := float64(0.0)
					score += organism.State.Maturity
					score += organism.State.Energy

					if score < rand.Float64() {
						// flee
						update.State.Position = position.Sub(velocity).Add(noise)
					} else {
						// seek
						update.State.Position = position.Add(velocity).Add(noise)
					}
				}
				clampToBounds(&update.State.Position)
				return
			}
		}

		// wander aimlessly
		position := organism.State.Position
		speed := float32(organism.Movement())
		noise := util.RandomDirection().Mul(speed)
		update.State.Position = position.Add(noise)
		clampToBounds(&update.State.Position)
	}
}
