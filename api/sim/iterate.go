package sim

// Iterate applies one iteration of AI and returns the change in state as
// a map of changes.
func Iterate(organisms map[string]*Organism) map[string]*Update {
	updates := make(map[string]*Update)

	// ensure an update is available for all orgs
	for _, organism := range organisms {
		// create update
		update := &Update{
			ID: organism.ID,
			State: &State{
				Type:     "alive",
				Energy:   organism.State.Energy,
				Position: organism.State.Position,
				Maturity: organism.State.Maturity,
			},
		}
		updates[organism.ID] = update
	}

	// process all upates
	for _, organism := range organisms {
		// get update
		update, _ := updates[organism.ID]

		// apply ai
		ApplyAI(update, updates, organism, PerceptionTest(organism, organisms))

		// apply constraints
		ApplyConstraints(update)
	}

	return updates
}
