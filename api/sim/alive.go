package sim

import (
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
	log "github.com/unchartedsoftware/plog"
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

		if len(perception.DistancePairs) > 0 {

			// Determine target
			bestScore := 0.0;
			targetPair := perception.DistancePairs[0];
			targetOrganism := perception.Organisms[0];

			log.Info("DistancePairs", perception.DistancePairs)
			for i, pair := range perception.DistancePairs {
				score := 1.0;

				if perception.Organisms[i] != nil {
					score += 1;
				}
				if (perception.Positions[i] != nil) {
					score += 1;
				}
				score += 1 - targetPair.Distance

				if score > bestScore {
					targetPair = pair
					targetOrganism = perception.Organisms[i]
					bestScore = score
				}
			}

			if targetPair != nil {
				runAway := false
				if targetOrganism != nil {
					// Check if we should be running away
					if targetPair.Organism.State.Energy > organism.State.Energy {
						runAway = true;
					}
				}
				if runAway {
					// move away from
					dir := organism.State.Position.Sub(targetPair.Organism.State.Position).Normalize()
					update.State.Position = organism.State.Position.Add(dir.Mul(float32(organism.Attributes.Speed)))
				} else {
					// Chase after
					dir := organism.State.Position.Add(targetPair.Organism.State.Position).Normalize()
					update.State.Position = organism.State.Position.Add(dir.Mul(float32(organism.Attributes.Speed)))
				}
			}
		}
	}
}
