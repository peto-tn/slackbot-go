package slackbot

import (
	"net/http"
)

// Start the http server. use net/http
func ListenAndServe(pattern, addr string, handler http.Handler) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		OnCall(w, r)
	})
	http.ListenAndServe(addr, handler)
}
