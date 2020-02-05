package sqsclient

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSClient struct {
	queueURL string
}

func New(queueURL string) *SQSClient {
	return &SQSClient{
		queueURL: queueURL,
	}
}

func SendMessage(message *sqs.SendMessageOutput) bool {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)
}

func ReceiveMessage(message *sqs.SendMessageOutput) bool {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)
}

func DeleteMessage(message *sqs.SendMessageOutput) bool {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)
}
