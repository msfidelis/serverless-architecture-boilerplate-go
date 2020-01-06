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

	"serverless-architecture-boilerplate-go/pkg/book"
	"serverless-architecture-boilerplate-go/pkg/dynamodb"
)

type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	dynamoTable := os.Getenv("DYNAMO_TABLE_BOOKS")
	client := dynamodb.New(dynamoTable)

	id, _ := uuid.NewUUID()

	book := &book.Book{
		Hashkey: id.String(),
		Created: time.Now().String(),
		Updated: false,
	}

	json.Unmarshal([]byte(request.Body), book)

	client.Save(book)

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
