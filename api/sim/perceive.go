package sim

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type OrganismPair struct {
	Distance float64
	Organism *Organism
}

type PositionPair struct {
	Distance float64
	Position mgl32.Vec3
}

type PerceptionResults struct {
	Organisms     []*OrganismPair
	Positions     []*PositionPair
	Directions    []mgl32.Vec3
}

// PerceptionTest results from an organisms perception test
func PerceptionTest(organism *Organism, targets map[string]*Organism) *PerceptionResults {
	organisms := make([]*OrganismPair, 0)
	positions := make([]*PositionPair, 0)
	directions := make([]mgl32.Vec3, 0)

	for _, target := range targets {
		if target.ID == organism.ID || target.State.Type == "dead" {
			continue
		}
		diff := target.State.Position.Sub(organism.State.Position)
		dir := diff.Normalize()
		dist := float64(diff.Len())
		// take sizes into account
		dist = math.Max(0.0, dist-target.State.Size-organism.State.Size)
		if dist <= organism.Attributes.Perception {
			organisms = append(organisms, &OrganismPair{
				Distance: dist,
				Organism: target,
			})
		}
		if dist <= organism.Attributes.Perception*2 {
			positions = append(positions, &PositionPair{
				Distance: dist,
				Position: target.State.Position,
			})
		}
		directions = append(directions, dir)
	}
	return &PerceptionResults{
		Organisms:  organisms,
		Positions:  positions,
		Directions: directions,
	}
}
