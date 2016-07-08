package sim

// Iterate applies one iteration of AI and returns the change in state as
// a map of changes.
func Iterate(organisms map[string]*Organism, delta int64) map[string]*Update {
	updates := make(map[string]*Update)

	// ensure an update is available for all orgs
	for _, organism := range organisms {
		// create update
		update := &Update{
			ID: organism.ID,
			State: &State{
				Type:     "alive",
				Energy:   organism.State.Energy,
				Size:     organism.State.Size,
				Position: organism.State.Position,
				Rotation: organism.State.Rotation,
				Maturity: organism.State.Maturity,
			},
		}
		updates[organism.ID] = update
	}

	// process all upates
	for _, organism := range organisms {

		update, _ := updates[organism.ID]

		// apply constraints
		ApplyConstraints(update, organism, delta)

		// apply ai
		ApplyAI(update, updates, organism, PerceptionTest(organism, organisms))

		// handle events
		//HandleEvents();

		updates[organism.ID] = update

	}

	return updates
}
