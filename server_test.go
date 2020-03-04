package slackbot

import (
	"net/http"
	"testing"
)

func TestListenAndServe(t *testing.T) {
	go ListenAndServe("/", ":8585", nil)

	resp, err := http.Get("http://localhost:8585")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
}
