package dynamodb

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"fmt"
	"os"
)

type DynamoDbClient struct {
	tableName string
}

func New(tableName string) *DynamoDbClient {
	return &DynamoDbClient{
		tableName: tableName,
	}
}

func (d *DynamoDbClient) Save(item interface{}) *dynamodb.PutItemOutput {

	hahaha, _ := json.Marshal(item)
	fmt.Println(string(hahaha))

	av, errMarsh := dynamodbattribute.MarshalMap(item)

	fmt.Println(av)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	if errMarsh != nil {
		fmt.Println("Error to marshalling new item:")
		fmt.Println(errMarsh.Error())
		os.Exit(1)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(d.tableName),
	}

	response, errPut := svc.PutItem(input)

	if errPut != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(errPut.Error())
		os.Exit(1)
	}

	return response
}

func (d DynamoDbClient) Find() string {
	return "find"
}

func (d DynamoDbClient) Query() string {
	return "query"
}

func (d DynamoDbClient) Scan() string {
	return "scan"
}

func (d DynamoDbClient) Update() string {
	return "updated"
}

func (d DynamoDbClient) UpdateItem() string {
	return "updated"
}

func (d DynamoDbClient) RemoveItem() string {
	return "removed"
}
