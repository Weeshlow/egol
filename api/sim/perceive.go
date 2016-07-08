package sim

type OrganismPair struct {
	Distance float64
	Organism *Organism
}

type PerceptionResults struct {
	Threats []*OrganismPair
	Family  []*OrganismPair
}

// PerceptionTest results from an organisms perception test
func PerceptionTest(organism *Organism, targets map[string]*Organism) *PerceptionResults {
	threats := make([]*OrganismPair, 0)
	family := make([]*OrganismPair, 0)
	for _, target := range targets {
		if target.ID == organism.ID || target.State.Type == "dead" {
			continue
		}
		dist, success := organism.Perceive(target)
		if success {
			if target.Attributes.Family == organism.Attributes.Family {
				family = append(family, &OrganismPair{
					Distance: dist,
					Organism: target,
				})
			} else {
				threats = append(threats, &OrganismPair{
					Distance: dist,
					Organism: target,
				})
			}

		}
	}
	return &PerceptionResults{
		Threats: threats,
		Family:  family,
	}
}
