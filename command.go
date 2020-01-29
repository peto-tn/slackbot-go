package slackbot

import (
	"errors"
	"reflect"
	"strings"
)

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

func SetupCommand(custom []*Command) {
	AddCommand(helpCommand)
	AddCommand(pingCommand)

	for _, c := range custom {
		AddCommand(c)
	}
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

func AddCommand(c *Command) {
	commands[c.Name] = c
	commandKeys = append(commandKeys, c.Name)
}

func Help(c *Command, desc bool) string {
	name := c.Name
	message := c.HelpMessage

	optionObj := c.Option
	option := ""
	if optionObj != nil {
		rt := reflect.TypeOf(optionObj)
		for i := 0; i < rt.NumField(); i++ {
			f := rt.Field(i)
			if option != "" {
				option += " "
			}
			option += "["
			option += f.Name

			optionValue := ""
			defaultValue := f.Tag.Get("default")

			if desc {
				choiceValue := f.Tag.Get("choice")
				if choiceValue != "" {
					if defaultValue != "" {
						optionValue = strings.Replace(choiceValue, defaultValue, "*"+defaultValue+"*", 1)
					} else {
						optionValue = choiceValue
					}
				}
			}
			if optionValue == "" && defaultValue != "" {
				optionValue = "*" + defaultValue + "*"
			}

			if optionValue != "" {
				option += "(" + optionValue + ")"
			}
			option += "]"
		}
	}

	help := name
	if option != "" {
		help += " " + option
	}
	if desc && message != "" {
		help += " : " + message
	}

	return help
}

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
		if choiceValue != "" {
			tmpValue := value
			value = ""
			for _, choice := range strings.Split(choiceValue, ",") {
				if choice == tmpValue {
					value = tmpValue
					break
				}
			}
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
			help += Help(commands[key], option.Desc == "true") + "\n"
		}
		PostEphemeral(e, help)
	},
	Option: HelpCommandOption{},
}

type HelpCommandOption struct {
	Desc string `default:"false" choice:"false,true"`
}

// PingCommand
var pingCommand = &Command{
	Name:        "ping",
	HelpMessage: "Reply pong.",

	Execute: func(e Event, opt interface{}) {
		ReplyMessage(e, "pong! :table_tennis_paddle_and_ball:")
	},
}
