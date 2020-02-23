package main

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"

	"serverless-architecture-boilerplate-go/pkg/libs/dynamoclient"
	"serverless-architecture-boilerplate-go/pkg/libs/sqsclient"
	"serverless-architecture-boilerplate-go/pkg/models/book"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	dynamoTable := os.Getenv("DYNAMO_TABLE_BOOKS")
	client := dynamoclient.New(dynamoTable)

	sqsQueue := os.Getenv("SQS_QUEUE_BOOKS")
	sqs := sqsclient.New(sqsQueue)

	id, _ := uuid.NewUUID()

	book := &book.Book{
		Hashkey:   id.String(),
		Created:   time.Now().String(),
		Updated:   time.Now().String(),
		Processed: false,
	}

	json.Unmarshal([]byte(request.Body), book)

	client.Save(book)
	sqs.SendMessage(book)

	body, err := json.Marshal(book)

	if err != nil {
		return Response{StatusCode: 404}, err
	}

	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      201,
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
