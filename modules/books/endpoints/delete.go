package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer
	hashkey := request.PathParameters["hashkey"]
	dynamoTable := os.Getenv("DYNAMO_TABLE_BOOKS")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"hashkey": {
				S: aws.String(hashkey),
			},
		},
		TableName: aws.String(dynamoTable),
	}

	_, errDelete := svc.DeleteItem(input)
	if errDelete != nil {
		fmt.Println("Got error calling DeleteItem")
		fmt.Println(errDelete.Error())
		return Response{StatusCode: 404}, errDelete
	}

	body, err := json.Marshal(map[string]interface{}{
		"hashkey": hashkey,
		"status":  "deleted",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
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
