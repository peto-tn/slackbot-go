package slackbot

import (
	"github.com/slack-go/slack"
)

// AppHomeHandler for Slack
type AppHomeHandler interface {
	OnOpend(e Event) slack.Blocks
}

var (
	appHomeHandler AppHomeHandler
)

// SetAppHomeHandler for slackbot
func SetAppHomeHandler(handler AppHomeHandler) {
	appHomeHandler = handler
}

// onAppHomeOpend event
func onAppHomeOpend(e Event) {
	if appHomeHandler == nil {
		return
	}

	api.PublishView(
		e.User(),
		slack.HomeTabViewRequest{
			Type:   "home",
			Blocks: appHomeHandler.OnOpend(e),
		},
		"",
	)
}
