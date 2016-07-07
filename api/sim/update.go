package sim

import (
	"encoding/json"
)

// Update reprsents a single iterations update.
type Update struct {
	OrganismID uint32
	State      State
}

// Marshal returns the byte representation of an update.
func (u *Update) Marshal() ([]byte, error) {
	return json.Marshal(u)
}
