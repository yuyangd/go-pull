package receiver

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSHandler struct {
	Service SQSIface
}

type SQSIface interface {
}

func sqsClient() *sqs.SQS {
	// Create Session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create a S3 client from just a session.
	return sqs.New(sess)
}
