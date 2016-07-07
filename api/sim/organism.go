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
	// physical traits
	Size float64 `json:"size"`
	// position / orientation
	Position mgl32.Vec3 `json:"position"`
	Rotation float64    `json:"rotation"`
	// energy
	Energy float64 `json:"energy"`
	// attacking / defending / consuming
	// Target uint32 `json:"target,omitempty"`
}

// Attributes represents the attributes of an organism.
type Attributes struct {
	Family         uint32  `json:"family"`
	Offense        float64 `json:"offense"`
	Defense        float64 `json:"defense"`
	Agility        float64 `json:"agility"`
	Reproductivity float64 `json:"reproductivity"`
	OffspringSize  float64 `json:"offspringSize"`
	// coordinate based
	Range      float64 `json:"range"`
	Perception float64 `json:"perception"`
	Speed      float64 `json:"speed"`
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
			Type: "alive",
			Position: mgl32.Vec3{
				rand.Float32(),
				rand.Float32(),
				rand.Float32(),
			},
			Rotation: 0.0,
			Energy:   1.0,
			Size:     baseAttributes.OffspringSize,
		},
		Attributes: &Attributes{
			Family:         baseAttributes.Family,
			Offense:        math.Max(0, baseAttributes.Offense+(rand.Float64()*10-5)),
			Defense:        math.Max(0, baseAttributes.Defense+(rand.Float64()*10-5)),
			Agility:        math.Max(0, baseAttributes.Agility+(rand.Float64()*10-5)),
			Reproductivity: math.Max(0, baseAttributes.Reproductivity+(rand.Float64()*0.02-0.01)),
			// coordniate based
			OffspringSize: math.Max(0, baseAttributes.OffspringSize+(rand.Float64()*0.02-0.01)),
			Speed:         math.Max(0, baseAttributes.Speed+(rand.Float64()*0.02-0.01)),
			Perception:    math.Max(0, baseAttributes.Perception+(rand.Float64()*0.02-0.01)),
			Range:         math.Max(0, baseAttributes.Range+(rand.Float64()*0.02-0.01)),
		},
	}
}

func (o *Organism) Update(update *Update) {
	// state
	o.State.Size = update.State.Size
	o.State.Position = update.State.Position
	o.State.Rotation = update.State.Rotation
	o.State.Energy = update.State.Energy
}

func SpawnChild(organism *Organism) Organism {
	offspring := NewOrganism(organism.Attributes)
	offspring.State.Position = organism.State.Position
	return offspring
}

// Marshal returns the byte representation of an organism.
func (o *Organism) Marshal() ([]byte, error) {
	return json.Marshal(o)
}
