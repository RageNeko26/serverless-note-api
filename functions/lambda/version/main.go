package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MsgResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	b, _ := json.MarshalIndent(MsgResponse{
		Status:  "success",
		Message: "Version 1.0.0",
	}, "", " ")

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		StatusCode: 200,
		Body:       string(b),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
