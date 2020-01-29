package slackbot

import (
	"context"

	"github.com/apex/gateway"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func AWSLambdaHandler(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r, err := gateway.NewRequest(ctx, e)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	w := gateway.NewResponse()
	OnCall(w, r)

	return w.End(), nil
}

func AWSLambdaStart() {
	lambda.Start(AWSLambdaHandler)
}
