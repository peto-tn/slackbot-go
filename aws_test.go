package slackbot

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestAWSLambdaHandler(t *testing.T) {
	testRun := ToolsCreateTestRun(ClearCommand, nil)

	testRun(t, "normal test", func(t *testing.T) {
		ctx := context.Background()
		e := events.APIGatewayProxyRequest{}

		res, err := AWSLambdaHandler(ctx, e)

		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	testRun(t, "error test", func(t *testing.T) {
		ctx := context.Background()
		e := events.APIGatewayProxyRequest{Path: ":foo"}

		res, err := AWSLambdaHandler(ctx, e)

		assert.Error(t, err)
		assert.NotNil(t, res)
	})
}

func TestAWSLambdaStart(t *testing.T) {
	go AWSLambdaStart()
}
