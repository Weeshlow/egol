package sim

import (
	"math/rand"
)

func offenseVsDefense(attacker *Organism, defender *Organism) bool {
	offense := attacker.Attributes.Offense + attacker.Attributes.Agility
	defense := defender.Attributes.Defense + defender.Attributes.Agility
	prob := 0.0
	if offense > defense {
		prob = offense / defense
	} else {
		prob = defense / offense
	}
	return prob > rand.Float64()
}

// Attack returns true if an organism successfully attacks another
func Attack(attacker *Organism, defender *Organism) bool {
	diff := defender.State.Position.Sub(attacker.State.Position)
	dist := float64(diff.Len())
	if dist > attacker.Attributes.Range {
		// out of range
		return false
	}
	// return success or failure
	return offenseVsDefense(attacker, defender)
}

// Consume returns the amount consumed.
func Consume(attacker *Organism, defender *Organism) float64 {
	return rand.Float64() * 0.25
}
