package main

import (
	"encoding/json"
	"main/model"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type BodyPayload struct {
	NoteID  int    `json:"noteid"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Body    string `json:"body"`
	Created string `json:"created"`
}

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var body BodyPayload

	json.Unmarshal([]byte(request.Body), &body)

	b, _ := json.MarshalIndent(model.WebResponse{
		Message: "Successfully add note",
		Status:  "success",
		Data: map[string]interface{}{
			"noteid": body.NoteID,
			"title":  body.Title,
		},
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
