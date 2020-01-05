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

func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer

	client := dynamodb.New("dev-serverless-go-books-catalog")

	proj := expression.NamesList(expression.Name("hashkey"), expression.Name("title"), expression.Name("author"), expression.Name("price"))
	expr, errBuilder := expression.NewBuilder().WithProjection(proj).Build()

	if errBuilder != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(errBuilder.Error())
		os.Exit(1)
	}

	var books []book.Book

	result := client.Scan(expr)

	errUnmarsh := dynamodbattribute.UnmarshalListOfMaps(result.Items, &books)

	if errUnmarsh != nil {
		return Response{StatusCode: 500}, errUnmarsh
	}

	body, err := json.Marshal(books)

	if err != nil {
		return Response{StatusCode: 500}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
