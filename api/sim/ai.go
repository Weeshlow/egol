package sim

import (
	log "github.com/unchartedsoftware/plog"
)

// Returns results from an organisms perception test
func ApplyAI(update *Update, updates map[string]*Update, organism *Organism, perception *PerceptionResults) {
	log.Info(organism.State.Type)
	switch organism.State.Type {
	case "alive":
		AliveAI(update, updates, organism, perception)
		break
	case "reproducing":
		ReproduceAI(update, updates, organism, perception)
		break
	case "attack":
		AttackAI(update, updates, organism, perception)
		break
	case "dead":
		DeadAI(update, updates, organism, perception)
		break
	}
}
