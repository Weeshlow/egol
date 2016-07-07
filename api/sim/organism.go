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
	// health
	Hunger float64 `json:"hunger"`
	Energy float64 `json:"energy"`
	// attacking / defending / consuming
	Target uint32 `json:"target,omitempty"`
}

// Attributes represents the attributes of an organism.
type Attributes struct {
	Family         uint32  `json:"family"`
	Hunger         float64 `json:"hunger"`
	Energy         float64 `json:"energy"`
	Offense        float64 `json:"offense"`
	Defense        float64 `json:"defense"`
	Agility        float64 `json:"agility"`
	Range          float64 `json:"range"`
	Reproductivity float64 `json:"reproductivity"`
	Size           float64 `json:"size"`
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
			Hunger:   0.0,
			Energy:   1.0,
		},
		Attributes: &Attributes{
			Family:         baseAttributes.Family,
			Offense:        math.Max(0, baseAttributes.Offense+(rand.Float64()*10-5)),
			Defense:        math.Max(0, baseAttributes.Offense+(rand.Float64()*10-5)),
			Agility:        math.Max(0, baseAttributes.Offense+(rand.Float64()*10-5)),
			Range:          math.Max(0, baseAttributes.Offense+(rand.Float64()*10-5)),
			Reproductivity: math.Max(0, baseAttributes.Offense+(rand.Float64()*10-5)),
			Size:           math.Max(0, baseAttributes.Offense+(rand.Float64()*2-1)),
		},
	}
}

// Marshal returns the byte representation of an organism.
func (o *Organism) Marshal() ([]byte, error) {
	return json.Marshal(o)
}
