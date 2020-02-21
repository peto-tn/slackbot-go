package slackbot

import (
	"encoding/json"
	"io"
)

// Slack Event Payload
type Payload map[string]interface{}

// Decode json.
func DecodeJSON(r io.Reader) (Payload, error) {
	data := make(map[string]interface{})
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// Get information as String.
func (p Payload) String(key string) string {
	if v, ok := p[key]; !ok {
		return ""
	} else if vv, ok := v.(string); !ok {
		return ""
	} else {
		return vv
	}
}

// Get "type" information.
func (p Payload) Type() string {
	return p.String("type")
}

// Get "token" information.
func (p Payload) Token() string {
	return p.String("token")
}

// Get Event.
func (p Payload) Event() Event {
	return p["event"].(map[string]interface{})
}
