package togoist

import (
	"encoding/json"
	"github.com/satori/go.uuid"
)

type Command struct {
	Type   string                 `json:"type"`
	TempId uuid.UUID              `json:"temp_id"`
	UUID   uuid.UUID              `json:"uuid"`
	Args   map[string]interface{} `json:"args"`
}

// initialize and return a new Commnad struct
func NewCommand(cmdType string, args map[string]interface{}) *Command {
	return &Command{
		Type:   cmdType,
		TempId: uuid.Must(uuid.NewV4()),
		UUID:   uuid.Must(uuid.NewV4()),
		Args:   args,
	}
}

// convert a command struct to a string representation
func (cmd *Command) Stringify() string {
	str, _ := json.Marshal(cmd)
	return string(str)
}
