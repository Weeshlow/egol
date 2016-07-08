package sim

// ApplyAI applies the ai state.
func ApplyAI(update *Update, updates map[string]*Update, organism *Organism, perception *PerceptionResults) {
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
	}
}
