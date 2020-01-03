package main

import (
	"fmt"

	"serverless-architecture-boilerplate-go/pkg/dynamodb"

	"github.com/google/uuid"
)

func main() {

	type Book struct {
		Year   int
		Title  string
		Plot   string
		Rating float64
	}

	_, id := uuid.NewUUID()

	client1 := dynamodb.New("dev-serverless-go-books-catalog")
	fmt.Println(id)

	client2 := dynamodb.New("fodase")
	fmt.Println(client2.Save())
	fmt.Println(client1.Save())

}
