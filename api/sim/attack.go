package sim

import (
	"math"
	"math/rand"
)

// AttackAI processes the organism for the given state.
func AttackAI(update *Update, updates map[string]*Update, organism *Organism, perception *PerceptionResults) {
	if len(perception.Threats) > 0 {
		// organisms in sight
		for _, other := range perception.Threats {
			if organism.InRange(other.Distance, other.Organism) {
				// in attack range, lets attack
				dmg := organism.Attack() - other.Organism.Defense() + (0.1 * rand.Float64())
				dmg = math.Max(0.0, math.Min(1.0, dmg))
				updates[other.Organism.ID].State.Energy -= dmg
				update.State.Energy += dmg
				return
			}
		}
	}
	// return to idle state
	update.State.Type = "alive"
}
