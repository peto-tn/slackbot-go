slackbot-go
=======

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[license]: https://github.com/peto-tn/slackbot-go/blob/master/LICENSE

## Description
Chatbot for slack of golang.

## Setup Token and ID
You need to set up Token and ID in one of the following ways.

### Call Setup() function
Can be set by calling the Setup function
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

### Environment Variables
Automatically used if the following environment variables are set.
- SLACK_ACCESS_TOKEN
- SLACK_BOT_USER_ID
- SLACK_VERIFICATION_TOKEN

## Entry point
Create an entry point for `slackbot-go`.

### AWS Lambda
```
import (
    slackbot "github.com/peto-tn/slackbot-go"
)

func main() {
    slackbot.AWSLambdaStart()
}
```
### GCP Cloud Functions
```
import (
    "net/http"

    slackbot "github.com/peto-tn/slackbot-go"
)

func OnCall(w http.ResponseWriter, r *http.Request) {
    slackbot.OnCall(w, r)
}
```
### Standalone
```
import (
    slackbot "github.com/peto-tn/slackbot-go"
)

func main() {
    slackbot.ListenAndServe("/", ":8000", nil)
}
```

## Add ChatOps Command
WIP

## Author
[peto-tn](https://github.com/peto-tn)
