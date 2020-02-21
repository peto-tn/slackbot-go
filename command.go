package slackbot

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Command for Slack ChatOps.
type Command struct {
	Name        string
	HelpMessage string
	Execute     func(e Event, opt interface{})
	Option      interface{}
}

var (
	commands    = map[string]*Command{}
	commandKeys = []string{}
)

// SetupCommand for slackbot.
// help and ping command are added automativally.
func SetupCommand(custom []*Command) {
	ClearCommand()
	AddCommand(helpCommand)
	AddCommand(pingCommand)

	for _, c := range custom {
		AddCommand(c)
	}
}

// ClearCommand all for slackbot.
func ClearCommand() {
	commands = map[string]*Command{}
	commandKeys = []string{}
}

func executeCommand(e Event, texts []string) bool {
	if c, ok := commands[texts[0]]; ok {
		option, err := ParseOption(c, texts[1:])
		if err != nil {
			ReplyMessage(e, err.Error())
		} else {
			c.Execute(e, option)
		}

		return true
	}

	return false
}

// AddCommand for slackbot.
func AddCommand(c *Command) {
	commands[c.Name] = c
	commandKeys = append(commandKeys, c.Name)
}

// SetDefaultHelpDescription display.
func SetDefaultHelpDescription(description bool) {
	if description {
		helpCommand.Option = HelpCommandOptionDesc{}
	} else {
		helpCommand.Option = HelpCommandOptionSimple{}
	}
}

// Help message command.
func Help(c *Command, desc bool) string {
	name := c.Name
	message := selectString(desc && c.HelpMessage != "", c.HelpMessage, "")
	message = italicString(message)
	message = boldString(message)
	message = selectString(message != "", " : "+message, "")

	// return if option is null
	if c.Option == nil {
		return name + message
	}

	option := ""
	rt := reflect.TypeOf(c.Option)
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		defaultValue := f.Tag.Get("default")
		choiceValue := selectString(desc, f.Tag.Get("choice"), "")

		value := selectString(choiceValue != "", choiceValue, defaultValue)
		value = boldSubstring(value, defaultValue)
		value = addBrackets(value)

		option += fmt.Sprintf(" [%s%s]", f.Name, value)
	}

	return name + option + message
}

// ParseOption of the command.
func ParseOption(c *Command, options []string) (interface{}, error) {
	if c.Option == nil {
		return nil, nil
	}

	optionLen := len(options)

	rv := reflect.New(reflect.TypeOf(c.Option)).Elem()
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		value := ""
		if optionLen > i {
			value = options[i]
		} else {
			value = rt.Field(i).Tag.Get("default")
		}

		// validate by choice value
		choiceValue := rt.Field(i).Tag.Get("choice")
		if choiceValue != "" && !containsString(strings.Split(choiceValue, ","), value) {
			value = ""
		}

		if value == "" {
			return nil, errors.New("option error.\n" + Help(c, true))
		}

		rv.Field(i).SetString(value)
	}

	return rv.Interface(), nil
}

// HelpCommand
var helpCommand = &Command{
	Name:        "help",
	HelpMessage: "Displays all of the help commands.",

	Execute: func(e Event, opt interface{}) {
		option := opt.(HelpCommandOption)

		help := ""
		for _, key := range commandKeys {
			help += Help(commands[key], option.IsDescription() == "true") + "\n"
		}
		PostEphemeral(e, help)
	},
	Option: HelpCommandOptionDesc{},
}

// HelpCommandOption
type HelpCommandOption interface {
	IsDescription() string
}

// HelpCommandOptionDesc is Help Command Option with default description enabled.
type HelpCommandOptionDesc struct {
	Description string `default:"true" choice:"false,true"`
}

// IsDescription
func (o HelpCommandOptionDesc) IsDescription() string {
	return o.Description
}

// HelpCommandOptionSimple is Help Command Option with default description enabled.
type HelpCommandOptionSimple struct {
	Description string `default:"false" choice:"false,true"`
}

// IsDescription
func (o HelpCommandOptionSimple) IsDescription() string {
	return o.Description
}

// PingCommand
var pingCommand = &Command{
	Name:        "ping",
	HelpMessage: "Reply pong.",

	Execute: func(e Event, opt interface{}) {
		ReplyMessage(e, "pong! :table_tennis_paddle_and_ball:")
	},
}
