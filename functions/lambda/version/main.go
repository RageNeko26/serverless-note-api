package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
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

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	resp := MsgResponse{
		Message: "Version 1.0.1",
		Status:  "success",
	}

	j, _ := json.MarshalIndent(&resp, "", " ")

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Version"))
	})
	lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
}
