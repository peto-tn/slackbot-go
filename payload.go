package slackbot

import (
	"encoding/json"
	"io"
)

// Payload for Slack Event
type Payload map[string]interface{}

// DecodeJSON data.
func DecodeJSON(r io.Reader) (Payload, error) {
	data := make(map[string]interface{})
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// String data in Payload.
func (p Payload) String(key string) string {
	if v, ok := p[key]; !ok {
		return ""
	} else if vv, ok := v.(string); !ok {
		return ""
	} else {
		return vv
	}
}

// Type of Payload.
func (p Payload) Type() string {
	return p.String("type")
}

// Token in Payload.
func (p Payload) Token() string {
	return p.String("token")
}

// Event in Payload.
func (p Payload) Event() Event {
	var e Event = p["event"].(map[string]interface{})
	e.ModifyText()
	return e
}
