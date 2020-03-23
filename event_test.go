package slackbot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvent_String(t *testing.T) {
	t.Parallel()
	event := Event{
		"str": "test",
		"int": 1,
	}

	t.Run("normal test", func(t *testing.T) {
		result := event.String("str")
		assert.Equal(t, "test", result)
	})

	t.Run("not found test", func(t *testing.T) {
		result := event.String("hoge")
		assert.Equal(t, "", result)
	})

	t.Run("invalid type test", func(t *testing.T) {
		result := event.String("int")
		assert.Equal(t, "", result)
	})
}

func TestEvent_Type(t *testing.T) {
	t.Parallel()
	event := Event{
		"type": "test",
	}

	t.Run("normal test", func(t *testing.T) {
		result := event.Type()
		assert.Equal(t, "test", result)
	})
}

func TestEvent_Text(t *testing.T) {
	t.Parallel()
	event := Event{
		"text": "test",
	}

	t.Run("normal test", func(t *testing.T) {
		result := event.Text()
		assert.Equal(t, "test", result)
	})
}

func TestEvent_Channel(t *testing.T) {
	t.Parallel()
	event := Event{
		"channel": "test",
	}

	t.Run("normal test", func(t *testing.T) {
		result := event.Channel()
		assert.Equal(t, "test", result)
	})
}

func TestEvent_User(t *testing.T) {
	t.Parallel()
	event := Event{
		"user": "test",
	}

	t.Run("normal test", func(t *testing.T) {
		result := event.User()
		assert.Equal(t, "test", result)
	})
}

func TestEvent_ThreadTimestamp(t *testing.T) {
	t.Parallel()
	t.Run("thread test", func(t *testing.T) {
		event := Event{
			"thread_ts": "test",
		}
		result := event.ThreadTimestamp()
		assert.Equal(t, "test", result)
	})

	t.Run("not thread test", func(t *testing.T) {
		event := Event{
			"event_ts": "test",
		}
		result := event.ThreadTimestamp()
		assert.Equal(t, "test", result)
	})
}

func TestEvent_ModifyText(t *testing.T) {
	t.Parallel()
	t.Run("normal test", func(t *testing.T) {
		event := Event{
			"text": "ho     ge ho  ge",
		}
		assert.Equal(t, "ho     ge ho  ge", event.Text())
		event.ModifyText()
		assert.Equal(t, "ho ge ho ge", event.Text())
	})
}
