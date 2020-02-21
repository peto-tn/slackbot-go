package main

import (
	slackbot "github.com/peto-tn/slackbot-go"
	_ "github.com/peto-tn/slackbot-go/example/command"
)

func main() {
	slackbot.AWSLambdaStart()
}
