package gcp

import (
	"net/http"

	slackbot "github.com/peto-tn/slackbot-go"
	_ "github.com/peto-tn/slackbot-go/example"
)

func OnCall(w http.ResponseWriter, r *http.Request) {
	slackbot.OnCall(w, r)
}
