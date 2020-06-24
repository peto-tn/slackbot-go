package slackbot

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
)

// MessageHandler for Slack
type MessageHandler interface {
	OnMessage(e Event, texts []string)
	OnMentionMessage(e Event, texts []string)
}

var (
	messageHandler MessageHandler
)

// SetMessageHandler for slackbot
func SetMessageHandler(handler MessageHandler) {
	messageHandler = handler
}

func onMessage(e Event) {
	texts := strings.Split(strings.TrimSpace(e.Text()), " ")
	if texts[0] == fmt.Sprintf("<@%s>", slackBotUserID) {
		texts = texts[1:]
		if !executeCommand(e, texts) && messageHandler != nil {
			messageHandler.OnMentionMessage(e, texts)
		}
	} else {
		if !executeCommand(e, texts) && messageHandler != nil {
			messageHandler.OnMessage(e, texts)
		}
	}
}

func onMentionMessage(e Event) {
	texts := strings.Split(strings.TrimSpace(e.Text()), " ")
	if !executeCommand(e, texts) && messageHandler != nil {
		messageHandler.OnMentionMessage(e, texts)
	}
}

// PostMessage to Slack.
func PostMessage(e Event, message string) {
	channel := e.Channel()
	api.PostMessage(
		channel,
		slack.MsgOptionText(message, true),
	)
}

// PostEphemeral message to Slack.
func PostEphemeral(e Event, message string) {
	channel := e.Channel()
	api.PostEphemeral(
		channel,
		e.User(),
		slack.MsgOptionText(message, true),
	)
}

// ReplyMessage to Slack.
func ReplyMessage(e Event, message string) {
	channel := e.Channel()
	threadTimestamp := e.ThreadTimestamp()
	api.PostMessage(
		channel,
		slack.MsgOptionTS(threadTimestamp),
		slack.MsgOptionText(message, true),
	)
}
