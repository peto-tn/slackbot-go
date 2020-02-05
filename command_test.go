package slackbot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Tools_ClearCommand() {
	commands = map[string]*Command{}
	commandKeys = []string{}
}

func Tools_InitCommand() {
	Tools_ClearCommand()
	SetupCommand([]*Command{})
}

func Tools_CreateTestRun(setup, tearDown func()) func(t *testing.T, testName string, f func(t *testing.T)) {
	return func(t *testing.T, testName string, f func(t *testing.T)) {
		t.Run(testName, func(t *testing.T) {
			setup()
			defer tearDown()
			f(t)
		})
	}
}

func TestSetupCommand(t *testing.T) {
	setup := Tools_ClearCommand
	tearDown := Tools_InitCommand
	testRun := Tools_CreateTestRun(setup, tearDown)

	testRun(t, "empty input test", func(t *testing.T) {
		SetupCommand([]*Command{})

		assert.Len(t, commands, 2)
		assert.Equal(t, "help", commands["help"].Name)
		assert.Equal(t, "ping", commands["ping"].Name)

		assert.Len(t, commandKeys, 2)
		assert.Equal(t, "help", commandKeys[0])
		assert.Equal(t, "ping", commandKeys[1])
	})

	testRun(t, "custom command input test", func(t *testing.T) {
		SetupCommand([]*Command{&Command{Name: "test"}})

		assert.Len(t, commands, 3)
		assert.Equal(t, "test", commands["test"].Name)

		assert.Len(t, commandKeys, 3)
		assert.Equal(t, "test", commandKeys[2])
	})
}

func TestExecuteCommand(t *testing.T) {
	setup := Tools_InitCommand
	tearDown := Tools_InitCommand
	testRun := Tools_CreateTestRun(setup, tearDown)

	testRun(t, "execute test", func(t *testing.T) {
		called := false
		AddCommand(&Command{
			Name: "test",
			Execute: func(e Event, opt interface{}) {
				called = true
			},
		})

		texts := []string{"test"}
		result := executeCommand(Event{}, texts)

		assert.True(t, result)
		assert.True(t, called)
	})

	testRun(t, "parse option error test", func(t *testing.T) {
		called := false
		AddCommand(&Command{
			Name: "test",
			Execute: func(e Event, opt interface{}) {
				called = true
			},
			Option: struct {
				Test string `default:"false" choice:"false,true"`
			}{},
		})

		defer func() {
			assert.False(t, called)
			recover()
		}()
		texts := []string{"test", "invalid_option"}
		executeCommand(Event{}, texts)
		assert.Fail(t, "Must not reach.")
	})

	testRun(t, "undefined command test", func(t *testing.T) {
		texts := []string{"test"}
		result := executeCommand(Event{}, texts)

		assert.False(t, result)
	})
}
