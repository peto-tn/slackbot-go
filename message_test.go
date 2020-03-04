package slackbot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestMessageHandler struct {
	OnMessaged        bool
	OnMentionMessaged bool
}

func (h *TestMessageHandler) Clear() {
	h.OnMessaged = false
	h.OnMentionMessaged = false
}

func (h *TestMessageHandler) OnMessage(e Event, texts []string) {
	h.OnMessaged = true
}

func (h *TestMessageHandler) OnMentionMessage(e Event, texts []string) {
	h.OnMentionMessaged = true
}

func TestSetMessageHandler(t *testing.T) {
	handler := &TestMessageHandler{}
	SetMessageHandler(handler)

	assert.Equal(t, handler, messageHandler)
}

func TestOnMessage(t *testing.T) {
	slackBotUserID = "bot"

	var called bool
	handler := &TestMessageHandler{}

	clear := func() {
		ClearCommand()
		called = false
		handler.Clear()
		SetMessageHandler(handler)
		AddCommand(&Command{
			Name: "test",
			Execute: func(e Event, opt interface{}) {
				called = true
			},
		})
	}
	testRun := ToolsCreateTestRun(clear, nil)

	testRun(t, "excute command test", func(t *testing.T) {
		event := Event{
			"text": "test",
		}

		onMessage(event)
		assert.True(t, called)
		assert.False(t, handler.OnMessaged)
	})

	testRun(t, "handle message test", func(t *testing.T) {
		event := Event{
			"text": "ignore",
		}

		onMessage(event)
		assert.False(t, called)
		assert.True(t, handler.OnMessaged)
	})

	testRun(t, "execute mention command test", func(t *testing.T) {
		event := Event{
			"text": "<@bot> test",
		}

		onMessage(event)
		assert.True(t, called)
		assert.False(t, handler.OnMessaged)
	})

	testRun(t, "error test", func(t *testing.T) {
		SetMessageHandler(nil)
		event := Event{
			"text": "ignore",
		}

		onMessage(event)
		assert.False(t, called)
		assert.False(t, handler.OnMessaged)
	})
}

func TestOnMentionMessage(t *testing.T) {
	slackBotUserID = "bot"

	var called bool
	handler := &TestMessageHandler{}

	clear := func() {
		ClearCommand()
		called = false
		handler.Clear()
		SetMessageHandler(handler)
		AddCommand(&Command{
			Name: "test",
			Execute: func(e Event, opt interface{}) {
				called = true
			},
		})
	}
	testRun := ToolsCreateTestRun(clear, nil)

	testRun(t, "execute mention command test", func(t *testing.T) {
		event := Event{
			"text": "<@bot> test",
		}

		onMentionMessage(event)
		assert.True(t, called)
		assert.False(t, handler.OnMentionMessaged)
	})

	testRun(t, "handle mention message test", func(t *testing.T) {
		event := Event{
			"text": "<@bot> ignore",
		}

		onMessage(event)
		assert.False(t, called)
		assert.True(t, handler.OnMentionMessaged)
	})

	testRun(t, "error test", func(t *testing.T) {
		SetMessageHandler(nil)
		event := Event{
			"text": "ignore",
		}

		onMessage(event)
		assert.False(t, called)
	})
}

func TestPostMessage(t *testing.T) {
	event := Event{}
	clear := func() {
		Setup("", "", "")
	}

	testRun := ToolsCreateTestRun(clear, clear)

	testRun(t, "normal test", func(t *testing.T) {
		PostMessage(event, "test")
	})

	testRun(t, "error test", func(t *testing.T) {
		api = nil
		defer func() {
			recover()
		}()
		PostMessage(event, "test")
		assert.Fail(t, "do not reached.")
	})
}

func TestPostEphemeral(t *testing.T) {
	event := Event{}
	clear := func() {
		Setup("", "", "")
	}

	testRun := ToolsCreateTestRun(clear, clear)

	testRun(t, "normal test", func(t *testing.T) {
		PostEphemeral(event, "test")
	})

	testRun(t, "error test", func(t *testing.T) {
		api = nil
		defer func() {
			recover()
		}()
		PostEphemeral(event, "test")
		assert.Fail(t, "do not reached.")
	})
}

func TestReplyMessage(t *testing.T) {
	event := Event{}
	clear := func() {
		Setup("", "", "")
	}

	testRun := ToolsCreateTestRun(clear, clear)

	testRun(t, "normal test", func(t *testing.T) {
		ReplyMessage(event, "test")
	})

	testRun(t, "error test", func(t *testing.T) {
		api = nil
		defer func() {
			recover()
		}()
		ReplyMessage(event, "test")
		assert.Fail(t, "do not reached.")
	})
}
