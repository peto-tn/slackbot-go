package slackbot

import (
	"net/http"
)

func ListenAndServe(pattern, addr string, handler http.Handler) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		OnCall(w, r)
	})
	http.ListenAndServe(addr, handler)
}
