package sim

// Iterate applies one iteration of AI and returns the change in state as
// a map of changes.
func Iterate(organisms map[string]*Organism) map[string]*Update {
	updates := make(map[string]*Update)
	for _, organism := range organisms {

		// create update
		update := &Update{
			ID:    organism.ID,
			State: &State{},
		}

		// apply constraints
		ApplyConstraints(update, organism)

		// apply ai
		ApplyAI(update, organism, PerceptionTest(organism, organisms))

		updates[organism.ID] = update

	}

	return updates
}
