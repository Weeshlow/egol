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
		len(perception.Threats) == 0 {
		// attempt to reproduce
		update.State.Type = "reproducing"
	} else {

		if len(perception.Threats) > 0 {
			// organisms in sight
			var closestThreat *Organism
			shortestDist := math.MaxFloat64

			for _, other := range perception.Threats {
				if organism.InRange(other.Distance, other.Organism) {
					// in attack range, lets attack
					update.State.Type = "attack"
					return
				}
				if other.Distance < shortestDist {
					//targetOrganism = other.Organism
					closestThreat = other.Organism
					shortestDist = other.Distance
				}
			}

			if closestThreat != nil {
				target := closestThreat.State.Position
				position := organism.State.Position
				diff := target.Sub(position)
				dist := diff.Len()
				if dist > 0.0 {
					dir := diff.Normalize()
					speed := float32(organism.Speed())
					velocity := dir.Mul(speed)
					noise := util.RandomDirection().Mul(0.005)

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

		if len(perception.Family) > 0 {
			position := organism.State.Position
			fam := perception.Family[0].Organism.State.Position
			diff := fam.Sub(position)
			dist := diff.Len()
			if dist > 0.0 {
				dir := diff.Normalize()
				speed := float32(organism.Speed())
				velocity := dir.Mul(speed)
				noise := util.RandomDirection().Mul(0.005)
				// move away from family, to spread out
				update.State.Position = position.Sub(velocity).Add(noise)
			}
			clampToBounds(&update.State.Position)
			return
		}

		// wander aimlessly
		position := organism.State.Position
		speed := float32(organism.Speed())
		noise := util.RandomDirection().Mul(speed)
		update.State.Position = position.Add(noise)
		clampToBounds(&update.State.Position)
	}
}
