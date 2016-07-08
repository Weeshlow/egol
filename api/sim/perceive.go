package sim

import (
	"math"
)

type OrganismPair struct {
	Distance float64
	Organism *Organism
}

type PerceptionResults struct {
	Organisms []*OrganismPair
}

// PerceptionTest results from an organisms perception test
func PerceptionTest(organism *Organism, targets map[string]*Organism) *PerceptionResults {
	organisms := make([]*OrganismPair, 0)

	for _, target := range targets {
		if target.ID == organism.ID || target.State.Type == "dead" {
			continue
		}
		diff := target.State.Position.Sub(organism.State.Position)
		dist := float64(diff.Len())
		// take sizes into account
		dist = math.Max(0.0, dist-target.State.Size-organism.State.Size)
		if dist <= organism.Attributes.Perception {
			organisms = append(organisms, &OrganismPair{
				Distance: dist,
				Organism: target,
			})
		}
	}
	return &PerceptionResults{
		Organisms: organisms,
	}
}
