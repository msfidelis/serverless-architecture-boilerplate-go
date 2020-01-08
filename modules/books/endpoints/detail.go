package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"serverless-architecture-boilerplate-go/pkg/book"
	"serverless-architecture-boilerplate-go/pkg/dynamodb"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer
	var books []book.Book

	hashkey := request.PathParameters["hashkey"]
	dynamoTable := os.Getenv("DYNAMO_TABLE_BOOKS")

	client := dynamodb.New(dynamoTable)

	proj := expression.NamesList(expression.Name("hashkey"), expression.Name("title"), expression.Name("author"), expression.Name("price"), expression.Name("updated"), expression.Name("created"))
	filt := expression.Name("hashkey").Equal(expression.Value(hashkey))
	expr, errBuilder := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()

	if errBuilder != nil {
		fmt.Println("Got error building expression:")
		return Response{StatusCode: 500}, errBuilder
	}

	result := client.Scan(expr)

	if len(result.Items) <= 0 {

		body, err := json.Marshal(map[string]interface{}{
			"hashkey": hashkey,
			"message": "not found",
		})

		if err != nil {
			return Response{StatusCode: 500}, err
		}

		return Response{
			StatusCode:      404,
			IsBase64Encoded: false,
			Body:            string(body),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil

	}

	errUnmarsh := dynamodbattribute.UnmarshalListOfMaps(result.Items, &books)

	if errUnmarsh != nil {
		return Response{StatusCode: 500}, errUnmarsh
	}

	body, err := json.Marshal(books[0])

	if err != nil {
		return Response{StatusCode: 500}, err
	}

	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
