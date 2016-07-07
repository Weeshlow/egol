package sim

import (
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
)


// RandomPosition returns a random vec3
func RandomPosition() mgl32.Vec3 {
	return mgl32.Vec3{
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
	}
}

// AliveAI processes the organism for the given state
func AliveAI(update *Update, updates map[string]*Update, organism *Organism, perception *PerceptionResults) {
	update.State.Position = RandomPosition()
	update.State.Type = "alive";
}
