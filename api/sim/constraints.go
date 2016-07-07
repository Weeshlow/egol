package sim

import (
	"math"
)

// ApplyConstraints updates attributes by one iteration.
func ApplyConstraints(organisms map[string]*Organism) {
	for _, organism := range organisms {
		updateHunger(organism);
		updateEnergy(organism);
		updateState(organism);
	}
}

func updateHunger(organism *Organism) {
	attributes := organism.Attributes;

	sizeFactor := attributes.Size * 0.001;

	attributes.Hunger += 0.01 + sizeFactor;
}

func updateEnergy(organism *Organism) {
	attributes := organism.Attributes;

	sizeFactor := attributes.Size * 0.001;

	attributes.Energy -= 0.01 + sizeFactor;

	if attributes.Energy < 0.7 && attributes.Hunger > 0 {
		consumedHunger := math.Min(0.01, attributes.Hunger)
		attributes.Hunger -= consumedHunger
		attributes.Energy += consumedHunger * 2
	}
}

func updateState(organism *Organism) {
	attributes := organism.Attributes;

	if attributes.Energy == 0 && attributes.Hunger == 1 {
		// organism state is dead
		organism.State.Type = "dead"
	}

}
