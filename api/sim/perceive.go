package sim

import (

	"github.com/go-gl/mathgl/mgl32"
)

type PerceptionResults struct {
	Organisms  []*Organism
	Positions  []*mgl32.Vec3
	Directions []*mgl32.Vec3
}

// PerceptionTest results from an organisms perception test
func PerceptionTest(organism *Organism, targets map[string]*Organism) *PerceptionResults {
	organisms := make([]*Organism, 0)
	positions := make([]*mgl32.Vec3, 0)
	directions := make([]*mgl32.Vec3, 0)
	for _, target := range targets {
		diff := target.State.Position.Sub(organism.State.Position)
		dist := float64(diff.Len())
		dir := diff.Normalize()
		if dist <= organism.Attributes.Perception {
			organisms = append(organisms, target)
		} else if (dist >= organism.Attributes.Perception * 2) {
			positions = append(positions, &target.State.Position)
		} else {
			directions = append(directions, &dir)
		}
	}
	return &PerceptionResults{
		Organisms: organisms,
		Positions: positions,
		Directions: directions,
	}
}
