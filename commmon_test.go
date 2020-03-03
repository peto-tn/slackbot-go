package slackbot

import (
	"testing"
)

func ToolsCreateTestRun(setup, tearDown func()) func(t *testing.T, testName string, f func(t *testing.T)) {
	return func(t *testing.T, testName string, f func(t *testing.T)) {
		t.Run(testName, func(t *testing.T) {
			if setup != nil {
				setup()
			}
			defer func() {
				if tearDown != nil {
					tearDown()
				}
			}()
			f(t)
		})
	}
}
