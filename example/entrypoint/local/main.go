package main

import (
	"os"

	slackbot "github.com/peto-tn/slackbot-go"
	_ "github.com/peto-tn/slackbot-go/example/command"
)

func main() {
	port := os.Getenv("PORT")
	slackbot.ListenAndServe("/", ":"+port, nil)
}
