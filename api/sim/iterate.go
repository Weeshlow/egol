package sim

import (
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/unchartedsoftware/plog"
)

// Iterate applies one iteration of AI
func Iterate(organisms []*Organism) []*Update {

	for _, organism := range organisms {
		log.Info("state", organism.State)
		log.Info("position", organism.Position)

		organism.Position = mgl32.Vec3{rand.Float32(), rand.Float32(), rand.Float32()}
	}

	return make([]*Update, 0)
}
