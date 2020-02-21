slackbot-go
=======

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[license]: https://github.com/peto-tn/slackbot-go/blob/master/LICENSE

## Description
Chatbot for slack of golang.

## Setup
### Token,  ID
You need to set up Token and ID in one of the following ways.
- Setup() function
- Environment Variables

#### Setup() function
Can be set by calling the Setup function.
```
import (
    slackbot "github.com/peto-tn/slackbot-go"
)

const (
    SLACK_ACCESS_TOKEN       string = "xoxb-0123456789-012345678901-ABCDEFGHIJKLMOPQRSTUVWXY"
    SLACK_BOT_USER_ID        string = "U01234567"
    SLACK_VERIFICATION_TOKEN string = "abcdefghijklmopqrstuvwxy"
)

func init() {
    slackbot.Setup(SLACK_ACCESS_TOKEN, SLACK_BOT_USER_ID, SLACK_VERIFICATION_TOKEN)
}
```

#### Environment Variables
Automatically used if the following environment variables are set.
- SLACK_ACCESS_TOKEN
- SLACK_BOT_USER_ID
- SLACK_VERIFICATION_TOKEN

### Entry point
Create an entry point for `slackbot-go`.

#### AWS Lambda
```
import (
    slackbot "github.com/peto-tn/slackbot-go"
)

func main() {
    slackbot.AWSLambdaStart()
}
```

#### GCP Cloud Functions
```
import (
    "net/http"

    slackbot "github.com/peto-tn/slackbot-go"
)

func OnCall(w http.ResponseWriter, r *http.Request) {
    slackbot.OnCall(w, r)
}
```

#### Listen 
```
import (
    slackbot "github.com/peto-tn/slackbot-go"
)

func main() {
    slackbot.ListenAndServe("/", ":8000", nil)
}
```

## Add ChatOps Command
This is a sample command to repeat a message.  
You can optionally specify the number of repetitions and the font.
```
package example

import (
	"strconv"

	slackbot "github.com/peto-tn/slackbot-go"
)

func init() {
	slackbot.AddCommand(&slackbot.Command{
		Name:        "repeat",
		HelpMessage: "Repeat input message.",
		Execute:     repeat,
		Option:      RepeatOption{},
	})
}

type RepeatOption struct {
	Message string
	Count   string `default:"1"`
	Font    string `default:"thin" choice:"thin,bold,italic"`
}

func repeat(e slackbot.Event, opt interface{}) {
	option := opt.(RepeatOption)

	message := ""

	// Add messages as many as Count
	count, err := strconv.Atoi(option.Count)
	if err != nil {
		slackbot.ReplyMessage(e, "error: Invalid format for 'Count' option.")
		return
	}
	for i := 0; i < count; i++ {
		message += option.Message
	}

	// Font
	switch option.Font {
	case "bold":
		message = "*" + message + "*"
	case "italic":
		message = "_" + message + "_"
	}

	slackbot.ReplyMessage(e, message)
}
```

## Author
[peto-tn](https://github.com/peto-tn)
