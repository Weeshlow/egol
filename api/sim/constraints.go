package sim

// ApplyConstraints updates attributes by one iteration.
func ApplyConstraints(organisms []*Organism) {
	for _, organism := range organisms {
		updateHunger(organism);
		updateEnergy(organism);
		updateState(organism);
	}
}

func updateHunger(organism *Organism) {
	//attributes := organism.Attributes;
}

func updateEnergy(organism *Organism) {
	//attributes := organism.Attributes;
}

func updateState(organism *Organism) {
	attributes := organism.Attributes;

	if attributes.Energy == 1 && attributes.Hunger == 1 {
		// organism state is dead
		organism.State.Type = "dead"
	}

}