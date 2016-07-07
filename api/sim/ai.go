package sim

// Returns results from an organisms perception test
func ApplyAI(update *Update, updates map[string]*Update, organism *Organism, perception *PerceptionResults) {
	switch (organism.State.Type) {
		case "alive":
			AliveAI(update, updates, organism, perception)
			break
		case "reproducing":
			ReproduceAI(update, updates, organism, perception)
			break
		case "dead":
			DeadAI(update, updates, organism, perception)
			break
	}
}
