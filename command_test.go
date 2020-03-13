package slackbot

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ToolsInitCommand() {
	ClearCommand()
	SetupCommand([]*Command{})
}

func TestSetupCommand(t *testing.T) {
	testRun := ToolsCreateTestRun(ClearCommand, ToolsInitCommand)

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
	testRun := ToolsCreateTestRun(ToolsInitCommand, ToolsInitCommand)

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
	})

	testRun(t, "undefined command test", func(t *testing.T) {
		texts := []string{"test"}
		result := executeCommand(Event{}, texts)

		assert.False(t, result)
	})
}

func TestAddCommand(t *testing.T) {
	testRun := ToolsCreateTestRun(ClearCommand, ToolsInitCommand)

	testRun(t, "add test", func(t *testing.T) {
		AddCommand(&Command{Name: "test"})
		assert.Len(t, commands, 1)
		assert.Equal(t, "test", commands["test"].Name)
	})
}

func TestSetDefaultHelpDescription(t *testing.T) {
	testRun := ToolsCreateTestRun(ToolsInitCommand, ToolsInitCommand)

	testRun(t, "desc test", func(t *testing.T) {
		SetDefaultHelpDescription(true)
	})

	testRun(t, "simple test", func(t *testing.T) {
		SetDefaultHelpDescription(false)
	})
}

func TestHelp(t *testing.T) {
	testRun := ToolsCreateTestRun(ToolsInitCommand, ToolsInitCommand)

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
			Desc  string `default:"true" choice:"false,true"`
			Bool  bool   `default:"true"`
			Int32 int32  `default:"1" min:"-20" max:"100"`
			Int   int    `default:"10" min:"-200" max:"1000"`
		}{}}
		help := Help(command, false)
		assert.Equal(t, "test [Desc(*true*)] [Bool(*true*)] [Int32(*1*)] [Int(*10*)]", help)
	})

	testRun(t, "option desc test", func(t *testing.T) {
		command := &Command{Name: "test", HelpMessage: "message", Option: struct {
			Desc  string `default:"true" choice:"false,true"`
			Bool  bool   `default:"true"`
			Int32 int32  `default:"1" min:"-20" max:"100"`
			Int   int    `default:"10" min:"-200" max:"1000"`
		}{}}
		help := Help(command, true)
		assert.Equal(t, "test [Desc(false,*true*)] [Bool(*true*,false)] [Int32(*1*,min:-20,max:100)] [Int(*10*,min:-200,max:1000)] : *_message_*", help)
	})
}

func TestParseOption(t *testing.T) {
	testRun := ToolsCreateTestRun(ToolsInitCommand, ToolsInitCommand)

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

func TestContainsChoice(t *testing.T) {
	testRun := ToolsCreateTestRun(ToolsInitCommand, ToolsInitCommand)

	type Test struct {
		Interface interface{}
		String    string `default:"true" choice:"hoge"`
		Bool      bool   `default:"true"`
		Int32     int32  `default:"1" min:"-20" max:"100"`
		Int       int    `default:"10" min:"-200" max:"1000"`
	}
	testOption := Test{}

	rv := reflect.New(reflect.TypeOf(testOption)).Elem()
	rt := rv.Type()

	testRun(t, "undefined test", func(t *testing.T) {
		result := containsChoice(rt.Field(0), "")
		assert.False(t, result)
	})

	testRun(t, "string test", func(t *testing.T) {
		result := containsChoice(rt.Field(1), "hoge")
		assert.True(t, result)
	})

	testRun(t, "bool test", func(t *testing.T) {
		result := containsChoice(rt.Field(2), "true")
		assert.True(t, result)
	})

	testRun(t, "int32 test", func(t *testing.T) {
		result := containsChoice(rt.Field(3), "1")
		assert.True(t, result)
	})

	testRun(t, "int32 over max test", func(t *testing.T) {
		result := containsChoice(rt.Field(3), "101")
		assert.False(t, result)
	})

	testRun(t, "int32 short min test", func(t *testing.T) {
		result := containsChoice(rt.Field(3), "-21")
		assert.False(t, result)
	})

	testRun(t, "int32 parse error test", func(t *testing.T) {
		result := containsChoice(rt.Field(3), "hoge")
		assert.False(t, result)
	})
}

func TestParseChoice(t *testing.T) {
	testRun := ToolsCreateTestRun(ToolsInitCommand, ToolsInitCommand)

	type Test struct {
		Interface interface{}
		String    string `choice:"hoge"`
		Bool      bool
		Int32     int32 `default:"1" min:"-20" max:"100"`
		Int       int   `default:"10" min:"-200" max:"1000"`
	}
	testOption := Test{}

	rv := reflect.New(reflect.TypeOf(testOption)).Elem()
	rt := rv.Type()

	testRun(t, "undefined test", func(t *testing.T) {
		choiceValue := parseChoice(rt.Field(0))
		assert.Len(t, choiceValue, 0)
	})

	testRun(t, "string test", func(t *testing.T) {
		choiceValue := parseChoice(rt.Field(1))
		assert.Len(t, choiceValue, 1)
		assert.Equal(t, "hoge", choiceValue[0])
	})

	testRun(t, "bool test", func(t *testing.T) {
		choiceValue := parseChoice(rt.Field(2))
		assert.Len(t, choiceValue, 2)
		assert.Equal(t, "true", choiceValue[0])
		assert.Equal(t, "false", choiceValue[1])
	})

	testRun(t, "int32 test", func(t *testing.T) {
		choiceValue := parseChoice(rt.Field(3))
		assert.Len(t, choiceValue, 3)
		assert.Equal(t, "1", choiceValue[0])
		assert.Equal(t, "min:-20", choiceValue[1])
		assert.Equal(t, "max:100", choiceValue[2])
	})

	testRun(t, "int test", func(t *testing.T) {
		choiceValue := parseChoice(rt.Field(4))
		assert.Len(t, choiceValue, 3)
		assert.Equal(t, "10", choiceValue[0])
		assert.Equal(t, "min:-200", choiceValue[1])
		assert.Equal(t, "max:1000", choiceValue[2])
	})
}

func TestSetValue(t *testing.T) {
	testRun := ToolsCreateTestRun(ToolsInitCommand, ToolsInitCommand)

	type Test struct {
		Interface interface{}
		String    string `choice:"hoge"`
		Bool      bool
		Int32     int32 `default:"1" min:"-20" max:"100"`
		Int       int   `default:"10" min:"-200" max:"1000"`
	}
	testOption := Test{}

	testRun(t, "undefined test", func(t *testing.T) {
		rv := reflect.New(reflect.TypeOf(testOption)).Elem()
		err := setValue(rv.Field(0), "")
		assert.NoError(t, err)
		assert.Equal(t, testOption.Interface, rv.Interface().(Test).Interface)
	})

	testRun(t, "string test", func(t *testing.T) {
		rv := reflect.New(reflect.TypeOf(testOption)).Elem()
		err := setValue(rv.Field(1), "fuga")
		assert.NoError(t, err)
		assert.Equal(t, "fuga", rv.Interface().(Test).String)
	})

	testRun(t, "bool test", func(t *testing.T) {
		rv := reflect.New(reflect.TypeOf(testOption)).Elem()
		err := setValue(rv.Field(2), "true")
		assert.NoError(t, err)
		assert.True(t, rv.Interface().(Test).Bool)
	})

	testRun(t, "bool error test", func(t *testing.T) {
		rv := reflect.New(reflect.TypeOf(testOption)).Elem()
		err := setValue(rv.Field(2), "hoge")
		assert.Error(t, err)
	})

	testRun(t, "int32 test", func(t *testing.T) {
		rv := reflect.New(reflect.TypeOf(testOption)).Elem()
		err := setValue(rv.Field(3), "1")
		assert.NoError(t, err)
		assert.Equal(t, int32(1), rv.Interface().(Test).Int32)
	})

	testRun(t, "int32 parse error test", func(t *testing.T) {
		rv := reflect.New(reflect.TypeOf(testOption)).Elem()
		err := setValue(rv.Field(3), "hoge")
		assert.Error(t, err)
	})

	testRun(t, "int test", func(t *testing.T) {
		rv := reflect.New(reflect.TypeOf(testOption)).Elem()
		err := setValue(rv.Field(4), "1")
		assert.NoError(t, err)
		assert.Equal(t, int(1), rv.Interface().(Test).Int)
	})
}

func TestHelpCommand_Execute(t *testing.T) {
	testRun := ToolsCreateTestRun(ToolsInitCommand, ToolsInitCommand)

	testRun(t, "normal test", func(t *testing.T) {
		helpCommand.Execute(Event{}, HelpCommandOptionDesc{})
	})
}

func TestHelpCommandOptionDesc_IsDescription(t *testing.T) {
	SetDefaultHelpDescription(true)
	t.Run("normal test", func(t *testing.T) {
		option, _ := ParseOption(helpCommand, []string{})
		result := option.(HelpCommandOption).IsDescription()
		assert.Equal(t, "true", result)
	})
}

func TestHelpCommandOptionSimple_IsDescription(t *testing.T) {
	SetDefaultHelpDescription(false)
	t.Run("normal test", func(t *testing.T) {
		option, _ := ParseOption(helpCommand, []string{})
		result := option.(HelpCommandOption).IsDescription()
		assert.Equal(t, "false", result)
	})
}

func TestPingCommand_Execute(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		pingCommand.Execute(Event{}, nil)
	})
}
