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
	Type     string     `json:"type"`
	Size     float64    `json:"size"`
	Position mgl32.Vec3 `json:"position"`
	Rotation float64    `json:"rotation"`
	Energy   float64    `json:"energy"`
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

func mutate(value, variance, min, max float64) float64 {
	mutation := (rand.Float64() * (variance * 2)) - variance
	return math.Min(max, math.Max(min, value+mutation))
}

func NewOrganism(baseAttributes *Attributes) *Organism {
	return &Organism{
		ID: util.RandID(),
		State: &State{
			Type: "alive",
			Position: mgl32.Vec3{
				rand.Float32(),
				rand.Float32(),
				rand.Float32(),
			},
			Rotation: 0.0,
			Energy:   0.8 + rand.Float64() * 0.2,
			Size:     mutate(baseAttributes.OffspringSize, 0.01, 0.1, 0.9),
		},
		Attributes: &Attributes{
			Family:         baseAttributes.Family,
			Offense:        mutate(baseAttributes.Offense, 5, 0, math.MaxFloat64),
			Defense:        mutate(baseAttributes.Defense, 5, 0, math.MaxFloat64),
			Agility:        mutate(baseAttributes.Agility, 5, 0, math.MaxFloat64),
			Reproductivity: mutate(baseAttributes.Reproductivity, 0.01, 0.1, 0.9),
			// coordniate based
			OffspringSize: mutate(baseAttributes.OffspringSize, 0.01, 0, 1.0),
			Speed:         mutate(baseAttributes.Speed, 0.01, 0, 1.0),
			Perception:    mutate(baseAttributes.Perception, 0.01, 0, 1.0),
			Range:         mutate(baseAttributes.Range, 0.01, 0, 1.0),
		},
	}
}

func (o *Organism) Update(update *Update) {
	o.State.Type = update.State.Type
	o.State.Size = update.State.Size
	o.State.Position = update.State.Position
	o.State.Rotation = update.State.Rotation
	o.State.Energy = update.State.Energy
}

func (o *Organism) Spawn() *Organism {
	offspring := NewOrganism(o.Attributes)
	offspring.State.Position = o.State.Position
	return offspring
}

// Marshal returns the byte representation of an organism.
func (o *Organism) Marshal() ([]byte, error) {
	return json.Marshal(o)
}
