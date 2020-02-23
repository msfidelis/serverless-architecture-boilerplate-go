package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"serverless-architecture-boilerplate-go/pkg/libs/dynamoclient"
	"serverless-architecture-boilerplate-go/pkg/models/book"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Response events.APIGatewayProxyResponse

func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer
	var filt expression.ConditionBuilder

	dynamoTable := os.Getenv("DYNAMO_TABLE_BOOKS")
	client := dynamoclient.New(dynamoTable)

	filterByProcessed := request.QueryStringParameters["processed"]
	fmt.Println(filterByProcessed)

	if filterByProcessed == "true" || filterByProcessed == "false" {
		b, err := strconv.ParseBool(filterByProcessed)
		if err != nil {
			fmt.Println("Got error building expression:")
			fmt.Println(err.Error())
		}
		filt = expression.Name("processed").Equal(expression.Value(b))
	} else {
		filt = expression.AttributeExists(expression.Name("hashkey"))
	}

	proj := expression.NamesList(expression.Name("hashkey"), expression.Name("title"), expression.Name("author"), expression.Name("price"), expression.Name("processed"))
	expr, errBuilder := expression.NewBuilder().WithProjection(proj).WithFilter(filt).Build()

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
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
