package sqsclient

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSClient struct {
	queueName string
}

func New(queueName string) *SQSClient {
	return &SQSClient{
		queueName: queueName,
	}
}

func (s *SQSClient) SendMessage(message interface{}) *sqs.SendMessageOutput {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	fmt.Println(s.queueName)

	svc := sqs.New(sess)

	resultURL, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(s.queueName),
	})
	if err != nil {
		fmt.Println("Error to get QueueURL:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	m, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error to marshall struct")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	result, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(m)),
		QueueUrl:    resultURL.QueueUrl,
	})
	if err != nil {
		fmt.Println("Error to get send message to sqs:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return result
}

func ReceiveMessage(message *sqs.SendMessageOutput) bool {
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))

	// svc := sqs.New(sess)

	return true
}

func DeleteMessage(message *sqs.SendMessageOutput) bool {
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))

	// svc := sqs.New(sess)

	return true
}
