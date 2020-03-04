package slackbot

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeJSON(t *testing.T) {
	t.Parallel()

	t.Run("normal test", func(t *testing.T) {
		payload, err := DecodeJSON(strings.NewReader(`{"test":"test"}`))
		assert.NoError(t, err)
		assert.NotNil(t, payload)
		assert.Equal(t, "test", payload.String("test"))
	})
}

func TestPayload_String(t *testing.T) {
	t.Parallel()
	payload := Payload{
		"str": "test",
		"int": 1,
	}

	t.Run("normal test", func(t *testing.T) {
		result := payload.String("str")
		assert.Equal(t, "test", result)
	})

	t.Run("not found test", func(t *testing.T) {
		result := payload.String("hoge")
		assert.Equal(t, "", result)
	})

	t.Run("invalid type test", func(t *testing.T) {
		result := payload.String("int")
		assert.Equal(t, "", result)
	})
}

func TestPayload_Type(t *testing.T) {
	t.Parallel()
	payload := Payload{
		"type": "test",
	}

	t.Run("normal test", func(t *testing.T) {
		result := payload.Type()
		assert.Equal(t, "test", result)
	})
}

func TestPayload_Token(t *testing.T) {
	t.Parallel()
	payload := Payload{
		"token": "test",
	}

	t.Run("normal test", func(t *testing.T) {
		result := payload.Token()
		assert.Equal(t, "test", result)
	})
}

func TestPayload_Event(t *testing.T) {
	t.Parallel()
	payload := Payload{
		"event": map[string]interface{}{
			"test": "test",
		},
	}

	t.Run("normal test", func(t *testing.T) {
		result := payload.Event()
		assert.NotNil(t, result)
		assert.Equal(t, "test", result.String("test"))
	})
}
