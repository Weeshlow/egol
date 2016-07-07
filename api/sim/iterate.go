package sim

import (
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/unchartedsoftware/plog"
)

// Iterate applies one iteration of AI and returns the change in state as
// a map of changes.
func Iterate(organisms map[string]*Organism) map[string]*Update {
	updates := make(map[string]*Update)
	for _, organism := range organisms {
		log.Info("state", organism.State)
		log.Info("position", organism.Position)
		updates[organism.ID] = &Update{
			ID: organism.ID,
			State: &State{
				Position: mgl32.Vec3{
					rand.Float32(),
					rand.Float32(),
					rand.Float32(),
				},
			},
		}
	}
	return updates
}
