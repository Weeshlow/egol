package sim

import (
	"encoding/json"
)

// Update reprsents a single iterations update.
type Update struct {
	ID         string      `json:"id"`
	State      *State      `json:"state"`
	Attributes *Attributes `json:"attributes"`
}

// Marshal returns the byte representation of an update.
func (u *Update) Marshal() ([]byte, error) {
	return json.Marshal(u)
}
