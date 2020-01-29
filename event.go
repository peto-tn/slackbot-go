package slackbot

type Event map[string]interface{}

func (e Event) String(key string) string {
	if v, ok := e[key]; !ok {
		return ""
	} else if vv, ok := v.(string); !ok {
		return ""
	} else {
		return vv
	}
}

func (e Event) Type() string {
	return e.String("type")
}

func (e Event) Text() string {
	return e.String("text")
}

func (e Event) Channel() string {
	return e.String("channel")
}

func (e Event) User() string {
	return e.String("user")
}

func (e Event) ThreadTimestamp() string {
	threadTimestamp := e.String("thread_ts")
	if threadTimestamp != "" {
		return threadTimestamp
	}

	return e.String("event_ts")
}
