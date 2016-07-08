package sim

// OrganismPair represents the preceived organism and the dsitance to it.
type OrganismPair struct {
	Distance float64
	Organism *Organism
}

// PerceptionResults represents the results from a perception check.
type PerceptionResults struct {
	Threats []*OrganismPair
	Family  []*OrganismPair
}

// PerceptionTest returns perception results for a given organism.
func PerceptionTest(organism *Organism, organisms map[string]*Organism) *PerceptionResults {
	threats := make([]*OrganismPair, 0)
	family := make([]*OrganismPair, 0)
	for _, target := range organisms {
		if target.ID == organism.ID {
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
