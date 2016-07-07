package sim

import (
	"github.com/go-gl/mathgl/mgl32"
)

func HighestPriorityTarget(organism *Organism, perception *PerceptionResults) *mgl32.Vec3 {
	var targetPosition *mgl32.Vec3
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
				targetPosition = &other.Organism.State.Position
				bestScore = score
			}
		}
		return targetPosition
	} else if len(perception.Positions) > 0 {
		// positions in sight
		bestScore := 0.0
		for _, other := range perception.Positions {
			score := 0.0
			score += 1 - other.Distance
			if score > bestScore {
				targetPosition = &other.Position
				bestScore = score
			}
		}
	}
	return targetPosition
}

// ReproduceAI processes the organism for the given state
func AttackAI(update *Update, updates map[string]*Update, organism *Organism, perception *PerceptionResults) {
	// get highest priority target
	target := HighestPriorityTarget(organism, perception)
	if target != nil {
		dist := float64(organism.State.Position.Sub(*target).Len())
		if dist <= organism.Attributes.Range {
			update.State.Position = *target
			return
		}
	}
	// return to idle state
	update.State.Type = "alive"
}
