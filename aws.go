package slackbot

import (
	"context"

	"github.com/apex/gateway"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// AWSLambdaHandler is handler when a slack event is received via aws lambda.
func AWSLambdaHandler(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r, err := gateway.NewRequest(ctx, e)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	w := gateway.NewResponse()
	OnCall(w, r)

	return w.End(), nil
}

// AWSLambdaStart is start execution of aws lambda.
func AWSLambdaStart() {
	lambda.Start(AWSLambdaHandler)
}
