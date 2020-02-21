package slackbot

// Slack Event
type Event map[string]interface{}

// Get information as String.
func (e Event) String(key string) string {
	if v, ok := e[key]; !ok {
		return ""
	} else if vv, ok := v.(string); !ok {
		return ""
	} else {
		return vv
	}
}

// Get "type" information.
func (e Event) Type() string {
	return e.String("type")
}

// Get "text" information.
func (e Event) Text() string {
	return e.String("text")
}

// Get "channel" information.
func (e Event) Channel() string {
	return e.String("channel")
}

// Get "user" information.
func (e Event) User() string {
	return e.String("user")
}

// Get event timestamp. If thread, get thread timestamp.
func (e Event) ThreadTimestamp() string {
	threadTimestamp := e.String("thread_ts")
	if threadTimestamp != "" {
		return threadTimestamp
	}

	return e.String("event_ts")
}
