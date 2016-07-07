package sim

import (
	"encoding/json"
	"math"
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/unchartedsoftware/egol/api/util"
)

// State represents the state of an organism.
type State struct {
	Type string `json:"type"`
	// position / orientation
	Position mgl32.Vec3 `json:"position,omitempty"`
	Rotation float32    `json:"rotation,omitempty"`
	// energy
	Energy float64 `json:"energy"`
	// attacking / defending / consuming
	Target uint32 `json:"target,omitempty"`
}

// Attributes represents the attributes of an organism.
type Attributes struct {
	Family         uint32  `json:"family"`
	Offense        float64 `json:"offense"`
	Defense        float64 `json:"defense"`
	Agility        float64 `json:"agility"`
	Range          float64 `json:"range"`
	Perception     float64 `json:"perception"`
	Reproductivity float64 `json:"reproductivity"`
	// calculate these based on above
	Speed float64 `json:"speed"`
	Size  float64 `json:"size"`
}

// Organism represents a single autonomous organism.
type Organism struct {
	ID         string      `json:"id"`
	State      *State      `json:"state"`
	Attributes *Attributes `json:"attributes"`
}

func NewOrganism(baseAttributes *Attributes) Organism {
	id := util.RandID()
	return Organism{
		ID: id,
		State: &State{
			Type:     "alive",
			Position: RandomPosition(),
			Energy:   1.0,
		},
		Attributes: &Attributes{
			Family:         baseAttributes.Family,
			Offense:        math.Max(0, baseAttributes.Offense+(rand.Float64()*10-5)),
			Defense:        math.Max(0, baseAttributes.Offense+(rand.Float64()*10-5)),
			Agility:        math.Max(0, baseAttributes.Offense+(rand.Float64()*10-5)),
			Reproductivity: math.Max(0, baseAttributes.Offense+(rand.Float64()*10-5)),
			// coordniate based
			Size:           math.Max(0, family.Size+(rand.Float64()*mutation)),
			Speed:          math.Max(0, family.Speed+(rand.Float64()*mutation)),
			Perception:     math.Max(0, family.Perception+(rand.Float64()*mutation)),
			Range:          math.Max(0, family.Range+(rand.Float64()*mutation)),
		},
	}
}

// Marshal returns the byte representation of an organism.
func (o *Organism) Marshal() ([]byte, error) {
	return json.Marshal(o)
}
