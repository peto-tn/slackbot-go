package main

import (
	"net/http"
	"os"

	slackbot "github.com/peto-tn/slackbot-go"
	_ "github.com/peto-tn/slackbot-go/example"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	slackbot.OnCall(w, r)
}
