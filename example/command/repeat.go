package command

import (
	"strconv"

	slackbot "github.com/peto-tn/slackbot-go"
)

func init() {
	slackbot.AddCommand(&slackbot.Command{
		Name:        "repeat",
		HelpMessage: "Repeat input message.",
		Execute:     repeat,
		Option:      RepeatOption{},
	})
}

type RepeatOption struct {
	Message string
	Count   string `default:"1"`
	Font    string `default:"thin" choice:"thin,bold,italic"`
}

func repeat(e slackbot.Event, opt interface{}) {
	option := opt.(RepeatOption)

	message := ""

	// Add messages as many as Count
	count, err := strconv.Atoi(option.Count)
	if err != nil {
		slackbot.ReplyMessage(e, "error: Invalid format for 'Count' option.")
		return
	}
	for i := 0; i < count; i++ {
		message += option.Message
	}

	// Font
	switch option.Font {
	case "bold":
		message = "*" + message + "*"
	case "italic":
		message = "_" + message + "_"
	}

	slackbot.ReplyMessage(e, message)
}
