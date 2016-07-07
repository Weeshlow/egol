package ws

// Message represents a basic message
type Message struct {
	Type    string      `json:"type"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// NewMessage returns a new message
func NewMessage(data []byte) (*Message, error) {
	return &Message{}, nil
}
