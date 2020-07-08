package slackbot

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

var (
	slackBotUserID    string
	verificationToken string
	accessToken       string

	api *slack.Client
)

// Setup slackbot.
func Setup(argBotUserID, argVerificationToken, argAccessToken string) {
	// get envrironment value
	if argBotUserID != "" {
		slackBotUserID = argBotUserID
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

// OnCall is receive slack events handler.
func OnCall(w http.ResponseWriter, r *http.Request) {
	if api == nil {
		Setup(
			os.Getenv("SLACK_BOT_USER_ID"),
			os.Getenv("SLACK_VERIFICATION_TOKEN"),
			os.Getenv("SLACK_ACCESS_TOKEN"),
		)
	}
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

		case "app_home_opened":
			verifyToken(w, p.Token())
			if verifyRequest(r) {
				// onAppHomeOpend
			}

		default:
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("not support event: %s", eventName)
			return
		}

	case "dialog_cancellation",
		"dialog_submission",
		"dialog_suggestion",
		"interactive_message",
		"message_action",
		"block_actions",
		"block_suggestion",
		"view_submission",
		"view_closed",
		"shortcut":
		var message slack.InteractionCallback
		if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
			log.Printf("unrecognized interactive message")
			return
		}
		onInteraction(message)

	default:
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("not support type: %s", typeName)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}

func verifyToken(w http.ResponseWriter, token string) {
	if token != verificationToken {
		w.WriteHeader(http.StatusInternalServerError)
		panic("not verified token: " + token)
	}
}

func verifyRequest(r *http.Request) bool {
	// ignore retry
	if _, ok := r.Header["X-Slack-Retry-Num"]; ok {
		return false
	}

	return true
}
