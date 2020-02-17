package slackbot

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/nlopes/slack"
)

var (
	slackBotUserId    string
	verificationToken string
	accessToken       string

	api *slack.Client
)

func init() {
	if flag.Lookup("test.v") != nil {
		return
	}

	Setup(
		os.Getenv("SLACK_BOT_USER_ID"),
		os.Getenv("SLACK_VERIFICATION_TOKEN"),
		os.Getenv("SLACK_ACCESS_TOKEN"),
	)
}

func Setup(argBotUserId, argVerificationToken, argAccessToken string) {
	// get envrironment value
	if argBotUserId != "" {
		slackBotUserId = argBotUserId
	}
	if argVerificationToken != "" {
		verificationToken = argVerificationToken
	}
	if argAccessToken != "" {
		accessToken = argAccessToken
	}

	// create slack client
	api = slack.New(accessToken)

	// setup default command
	SetupCommand([]*Command{})
}

func OnCall(w http.ResponseWriter, r *http.Request) {
	p, err := DecodeJSON(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	typeName := p.Type()
	switch typeName {
	case "url_verification":
		// verify url response
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(p.String("challenge")))
		return

	case "event_callback":
		event := p.Event()
		eventName := event.Type()
		switch eventName {
		case "message":
			verifyToken(w, p.Token())
			if verifyRequest(r) {
				onMessage(event)
			}

		case "app_mention":
			verifyToken(w, p.Token())
			if verifyRequest(r) {
				onMentionMessage(event)
			}

		default:
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalf("not support event: %s", eventName)
			return
		}

	default:
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("not support type: %s", typeName)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}

func verifyToken(w http.ResponseWriter, token string) {
	if token != verificationToken {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("not verified token: %s", token)
	}
}

func verifyRequest(r *http.Request) bool {
	// ignore retry
	if _, ok := r.Header["X-Slack-Retry-Num"]; ok {
		return false
	}

	return true
}
