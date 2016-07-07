package sim

import (
	"encoding/json"

	"github.com/go-gl/mathgl/mgl32"
)

// State represents the state of an organism.
type State struct {
	Type      uint32
	Timestamp uint64
}

// Attributes represents the attributes of an organism.
type Attributes struct {
	Family         uint32
	Hunger         float32
	Energy         float32
	Offense        uint32
	Defense        uint32
	Agility        uint32
	Range          float32
	Reproductivity uint32
}

// Organism represents a single autonomous organism.
type Organism struct {
	ID         uint32
	Position   mgl32.Vec3
	Rotation   float32
	State      *State
	Attributes *Attributes
}

// Marshal returns the byte representation of an organism.
func (o *Organism) Marshal() ([]byte, error) {
	return json.Marshal(o)
}
