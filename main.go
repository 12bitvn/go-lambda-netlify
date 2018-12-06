package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body, _ := json.Marshal(map[string]interface{}{
		"msg": "Hello from 12bit.vn",
	})
	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"content-type": "application/json",
		},
		StatusCode: 200,
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(handler)
}
