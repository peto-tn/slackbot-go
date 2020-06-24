package slackbot

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	clear := func() {
		ClearCommand()
		slackBotUserID = ""
		verificationToken = ""
		accessToken = ""
		api = nil
	}
	testRun := ToolsCreateTestRun(clear, nil)

	testRun(t, "normal test", func(t *testing.T) {
		Setup("test1", "test2", "test3")

		assert.Equal(t, "test1", slackBotUserID)
		assert.Equal(t, "test2", verificationToken)
		assert.Equal(t, "test3", accessToken)
		assert.NotNil(t, api)
	})

	testRun(t, "empty test", func(t *testing.T) {
		Setup("", "", "")

		assert.Equal(t, "", slackBotUserID)
		assert.Equal(t, "", verificationToken)
		assert.Equal(t, "", accessToken)
		assert.NotNil(t, api)
	})
}

func TestOnCall(t *testing.T) {
	clear := func() {
		ClearCommand()
		slackBotUserID = ""
		verificationToken = ""
		accessToken = ""
		api = nil
	}
	testRun := ToolsCreateTestRun(clear, nil)

	testRun(t, "create api test", func(t *testing.T) {
		os.Setenv("SLACK_BOT_USER_ID", "test1")
		os.Setenv("SLACK_VERIFICATION_TOKEN", "test2")
		os.Setenv("SLACK_ACCESS_TOKEN", "test3")

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "", strings.NewReader(""))

		OnCall(rec, req)

		assert.Equal(t, "test1", slackBotUserID)
		assert.Equal(t, "test2", verificationToken)
		assert.Equal(t, "test3", accessToken)
		assert.NotNil(t, api)
	})

	testRun(t, "url verification test", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "", strings.NewReader(`{"type":"url_verification", "challenge":"test"}`))

		OnCall(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "text/plain", rec.Header().Get("Content-Type"))
		assert.Equal(t, "test", rec.Body.String())
	})

	testRun(t, "message test", func(t *testing.T) {
		os.Setenv("SLACK_VERIFICATION_TOKEN", "token")
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "", strings.NewReader(`{"type":"event_callback", "token":"token", "event":{"type":"message"}}`))

		OnCall(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "text/plain", rec.Header().Get("Content-Type"))
	})

	testRun(t, "app mention test", func(t *testing.T) {
		os.Setenv("SLACK_BOT_USER_ID", "bot")
		os.Setenv("SLACK_VERIFICATION_TOKEN", "token")
		called := false
		AddCommand(&Command{
			Name: "test",
			Execute: func(e Event, opt interface{}) {
				called = true
			},
		})
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "", strings.NewReader(`{"type":"event_callback", "token":"token", "event":{"type":"app_mention", "text":"test"}}`))

		OnCall(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "text/plain", rec.Header().Get("Content-Type"))
		assert.True(t, called)
	})

	testRun(t, "not support event test", func(t *testing.T) {
		os.Setenv("SLACK_VERIFICATION_TOKEN", "token")
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "", strings.NewReader(`{"type":"event_callback", "token":"token", "event":{"type":"test"}}`))

		OnCall(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	testRun(t, "not support type test", func(t *testing.T) {
		os.Setenv("SLACK_VERIFICATION_TOKEN", "token")
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "", strings.NewReader(`{"type":"test"}}`))

		OnCall(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestVerifyToken(t *testing.T) {
	clear := func() {
		verificationToken = ""
	}
	testRun := ToolsCreateTestRun(clear, clear)

	testRun(t, "normal test", func(t *testing.T) {
		verificationToken = "test"

		rec := httptest.NewRecorder()

		verifyToken(rec, "test")

		assert.NotEqual(t, http.StatusInternalServerError, rec.Code)
	})

	testRun(t, "error test", func(t *testing.T) {
		verificationToken = "test"

		rec := httptest.NewRecorder()

		defer func() {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			recover()
		}()

		verifyToken(rec, "")
	})
}

func TestVerifyRequest(t *testing.T) {
	testRun := ToolsCreateTestRun(nil, nil)

	testRun(t, "normal test", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "", strings.NewReader(`{}`))

		result := verifyRequest(req)

		assert.True(t, result)
	})

	testRun(t, "error test", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "", strings.NewReader(`{}`))
		req.Header.Set("X-Slack-Retry-Num", "1")

		result := verifyRequest(req)

		assert.False(t, result)
	})
}
