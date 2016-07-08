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
	Position mgl32.Vec3 `json:"position"`
	Energy   float64    `json:"energy"`
	Maturity float64    `json:"maturity"`
}

// Attributes represents the attributes of an organism.
type Attributes struct {
	Family         uint32  `json:"family"`
	Offense        float64 `json:"offense"`
	Defense        float64 `json:"defense"`
	Agility        float64 `json:"agility"`
	Reproductivity float64 `json:"reproductivity"`
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
			Type:     "alive",
			Position: util.RandomPosition(),
			Maturity: 0.0,
			Energy:   0.5 + rand.Float64()*0.2 - 0.1,
		},
		Attributes: &Attributes{
			Family:         baseAttributes.Family,
			Offense:        mutate(baseAttributes.Offense, 0.01, 0, 1.0),
			Defense:        mutate(baseAttributes.Defense, 0.01, 0, 1.0),
			Reproductivity: mutate(baseAttributes.Reproductivity, 0.01, 0.1, 0.9),
			// coordniate based
			Perception: mutate(baseAttributes.Perception, 0.01, 0, 1.0),
			Range:      mutate(baseAttributes.Range, 0.01, 0, 1.0),
			Speed:      mutate(baseAttributes.Speed, 0.01, 0, 1.0),
		},
	}
}

func (o *Organism) Update(update *Update) {
	o.State.Type = update.State.Type
	o.State.Position = update.State.Position
	o.State.Energy = update.State.Energy
	o.State.Maturity = update.State.Maturity
}

func (o *Organism) Speed() float64 {
	return 0.01 + (o.Attributes.Speed * o.Size())
}

func (o *Organism) Size() float64 {
	return 0.005 + (o.State.Maturity * o.State.Energy * 0.01)
}

func (o *Organism) Attack() float64 {
	return o.Size() * o.Attributes.Offense
}

func (o *Organism) Defense() float64 {
	return o.Size() * o.Attributes.Defense
}

func (o *Organism) InRange(dist float64, other *Organism) bool {
	return (dist - other.Size() - o.Size()) < o.Attributes.Range
}

func (o *Organism) Perceive(other *Organism) (float64, bool) {
	diff := other.State.Position.Sub(o.State.Position)
	dist := float64(diff.Len())
	// take sizes into account
	if dist-other.Size()-o.Size() < o.Attributes.Perception {
		return dist, true
	}
	return 0, false
}

func (o *Organism) Spawn() *Organism {
	offspring := NewOrganism(o.Attributes)
	noise := util.RandomDirection().Mul(float32(rand.Float64() * o.Size()))
	offspring.State.Position = o.State.Position.Add(noise)
	return offspring
}

// Marshal returns the byte representation of an organism.
func (o *Organism) Marshal() ([]byte, error) {
	return json.Marshal(o)
}
