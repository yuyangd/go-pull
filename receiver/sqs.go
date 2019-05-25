package receiver

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSHandler struct {
	Service SQSIface
	SQSURL  *string
}

type SQSIface interface {
	ReceiveMessage(*sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error)
	DeleteMessage(*sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error)
}

func sqsClient() *sqs.SQS {
	// Create Session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create a S3 client from just a session.
	return sqs.New(sess)
}

func (h *SQSHandler) ReceiveMessage() {
	result, err := h.Service.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            h.SQSURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(20), // 20 seconds
		WaitTimeSeconds:     aws.Int64(0),
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	if len(result.Messages) == 0 {
		fmt.Println("Received no messages")
		return
	}

	// Use the message attributes to get the object
	fmt.Println("Message Key", *(result.Messages[0].MessageAttributes["Key"].StringValue))
	fmt.Println("Message Bucket", *(result.Messages[0].MessageAttributes["Bucket"].StringValue))
	// Download the object

	// Delete the message
	resultDelete, err := h.Service.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      h.SQSURL,
		ReceiptHandle: result.Messages[0].ReceiptHandle,
	})

	if err != nil {
		fmt.Println("Delete Error", err)
		return
	}

	fmt.Println("Message Deleted", *resultDelete)

	// Delete the object
}
