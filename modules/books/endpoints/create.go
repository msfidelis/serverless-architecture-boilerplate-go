package main

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"

	"serverless-architecture-boilerplate-go/pkg/dynamodb"

	"log"
)

type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	log.Print("Request body: ", request.Body)

	type Book struct {
		Hashkey      string      `json:"hashkey"`
		Title        string      `json:"title"`
		Author       string      `json:"author"`
		Price        float64     `json:"price"`
		Updated      bool        `json:"updated"`
		Created      string      `json:"created"`
		CustomStruct interface{} `json:",omitempty"`
	}

	id, _ := uuid.NewUUID()

	book := &Book{
		Hashkey: id.String(),
		Created: time.Now().String(),
		Updated: false,
	}

	json.Unmarshal([]byte(request.Body), book)

	// if errParseBody != nil {
	// http.Error(w, errParseBody.Error(), http.StatusBadRequest)
	// return
	// }

	// id, _ := uuid.NewUUID()

	// book := &Book{
	// 	Hashkey: id.String(),
	// 	Author:  "Neil Gaiman",
	// 	Price:   "20.00",
	// 	Updated: 0,
	// }

	client := dynamodb.New("dev-serverless-go-books-catalog")

	client.Save(book)

	body, err := json.Marshal(map[string]interface{}{
		"book": book,
	})

	if err != nil {
		return Response{StatusCode: 404}, err
	}

	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      201,
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
