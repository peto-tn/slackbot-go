package main

import (
	slackbot "github.com/peto-tn/slackbot-go"
	// add command
	_ "github.com/peto-tn/slackbot-go/example/command"
)

func main() {
	slackbot.AWSLambdaStart()
}
