package main

import (
	slackbot "github.com/peto-tn/slackbot-go"
	_ "github.com/peto-tn/slackbot-go/example"
)

func main() {
	slackbot.AWSLambdaStart()
}
