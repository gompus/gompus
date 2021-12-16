package commands

import (
	"encoding/json"
	"github.com/gompus/gompus/models/gateway/payloads"
)

func marshalCommand(op payloads.Op, cmd interface{}) ([]byte, error) {
	data, err := json.Marshal(cmd)
	if err != nil {
		return nil, err
	}

	fields := make(map[string]interface{})
	if err = json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}

	fields["op"] = op
	return json.Marshal(fields)
}
