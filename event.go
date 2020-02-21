package slackbot

// Event
type Event map[string]interface{}

// String data in Event.
func (e Event) String(key string) string {
	if v, ok := e[key]; !ok {
		return ""
	} else if vv, ok := v.(string); !ok {
		return ""
	} else {
		return vv
	}
}

// Type of Event.
func (e Event) Type() string {
	return e.String("type")
}

// Text of Event.
func (e Event) Text() string {
	return e.String("text")
}

// Channel of Event.
func (e Event) Channel() string {
	return e.String("channel")
}

// User of eVent.
func (e Event) User() string {
	return e.String("user")
}

// ThreadTimestamp of Event. If not thread, get event timestamp.
func (e Event) ThreadTimestamp() string {
	threadTimestamp := e.String("thread_ts")
	if threadTimestamp != "" {
		return threadTimestamp
	}

	return e.String("event_ts")
}
