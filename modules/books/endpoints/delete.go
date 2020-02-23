package main

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"serverless-architecture-boilerplate-go/pkg/libs/dynamoclient"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer
	var body []byte
	var statusCode int

	dynamoTable := os.Getenv("DYNAMO_TABLE_BOOKS")
	client := dynamoclient.New(dynamoTable)

	hashkey := request.PathParameters["hashkey"]

	key := map[string]*dynamodb.AttributeValue{
		"hashkey": {
			S: aws.String(hashkey),
		},
	}

	removed := client.RemoveItem(key)

	if removed == true {
		statusCode = 200
		payload, err := json.Marshal(map[string]interface{}{
			"hashkey": hashkey,
			"status":  "deleted",
		})
		if err == nil {
			body = payload
		}
	} else {
		statusCode = 404
		payload, err := json.Marshal(map[string]interface{}{
			"hashkey": hashkey,
			"status":  "not found",
		})
		if err != nil {
			body = payload
		}
	}

	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      statusCode,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
