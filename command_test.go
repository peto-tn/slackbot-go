package slackbot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Tools_InitCommand() {
	ClearCommand()
	SetupCommand([]*Command{})
}

func TestSetupCommand(t *testing.T) {
	testRun := Tools_CreateTestRun(ClearCommand, Tools_InitCommand)

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
	testRun := Tools_CreateTestRun(Tools_InitCommand, Tools_InitCommand)

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

func TestAddCommand(t *testing.T) {
	testRun := Tools_CreateTestRun(ClearCommand, Tools_InitCommand)

	testRun(t, "add test", func(t *testing.T) {
		AddCommand(&Command{Name: "test"})
		assert.Len(t, commands, 1)
		assert.Equal(t, "test", commands["test"].Name)
	})
}

func TestSetDefaultHelpDescription(t *testing.T) {
	testRun := Tools_CreateTestRun(Tools_InitCommand, Tools_InitCommand)

	testRun(t, "desc test", func(t *testing.T) {
		SetDefaultHelpDescription(true)
	})

	testRun(t, "simple test", func(t *testing.T) {
		SetDefaultHelpDescription(false)
	})
}

func TestHelp(t *testing.T) {
	testRun := Tools_CreateTestRun(Tools_InitCommand, Tools_InitCommand)

	testRun(t, "no message test", func(t *testing.T) {
		command := &Command{Name: "test"}
		help := Help(command, false)
		assert.Equal(t, "test", help)
	})

	testRun(t, "no message desc test", func(t *testing.T) {
		command := &Command{Name: "test"}
		help := Help(command, true)
		assert.Equal(t, "test", help)
	})

	testRun(t, "no option test", func(t *testing.T) {
		command := &Command{Name: "test", HelpMessage: "message"}
		help := Help(command, false)
		assert.Equal(t, "test", help)
	})

	testRun(t, "no option desc test", func(t *testing.T) {
		command := &Command{Name: "test", HelpMessage: "message"}
		help := Help(command, true)
		assert.Equal(t, "test : *_message_*", help)
	})

	testRun(t, "option test", func(t *testing.T) {
		command := &Command{Name: "test", HelpMessage: "message", Option: struct {
			Desc string `default:"true" choice:"false,true"`
		}{}}
		help := Help(command, false)
		assert.Equal(t, "test [Desc(*true*)]", help)
	})

	testRun(t, "option desc test", func(t *testing.T) {
		command := &Command{Name: "test", HelpMessage: "message", Option: struct {
			Desc string `default:"true" choice:"false,true"`
		}{}}
		help := Help(command, true)
		assert.Equal(t, "test [Desc(false,*true*)] : *_message_*", help)
	})
}

func TestParseOption(t *testing.T) {
	testRun := Tools_CreateTestRun(Tools_InitCommand, Tools_InitCommand)

	testRun(t, "no option test", func(t *testing.T) {
		command := &Command{Name: "test"}
		option, err := ParseOption(command, []string{})
		assert.NoError(t, err)
		assert.Nil(t, option)
	})

	testRun(t, "default test", func(t *testing.T) {
		type Test struct {
			Desc string `default:"true" choice:"false,true"`
		}
		expectOption := Test{Desc: "true"}
		command := &Command{Name: "test", Option: Test{}}
		option, err := ParseOption(command, []string{})
		assert.NoError(t, err)
		assert.Equal(t, expectOption, option)
	})

	testRun(t, "choice test", func(t *testing.T) {
		type Test struct {
			Desc string `default:"true" choice:"false,true"`
		}
		expectOption := Test{Desc: "false"}
		command := &Command{Name: "test", Option: Test{}}
		option, err := ParseOption(command, []string{"false"})
		assert.NoError(t, err)
		assert.Equal(t, expectOption, option)
	})

	testRun(t, "option error test", func(t *testing.T) {
		type Test struct {
			Desc string `default:"true" choice:"false,true"`
		}
		command := &Command{Name: "test", Option: Test{}}
		option, err := ParseOption(command, []string{"invalid"})
		assert.Error(t, err)
		assert.Nil(t, option)
	})
}
