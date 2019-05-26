package receiver

import (
	"errors"
	"log"

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

func SqsClient() *sqs.SQS {
	// Create Session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create a S3 client from just a session.
	return sqs.New(sess)
}

// ReceiveMessage receive message from FIFO queue
func (h *SQSHandler) ReceiveMessage() (*sqs.ReceiveMessageOutput, error) {
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
		return nil, err
	}

	if len(result.Messages) == 0 {
		return nil, errors.New("Received no messages")
	}

	return result, nil
	// Use the message attributes to get the object
	// fmt.Println("Message Key", *(result.Messages[0].MessageAttributes["Key"].StringValue))
	// fmt.Println("Message Bucket", *(result.Messages[0].MessageAttributes["Bucket"].StringValue))
}

// DeleteMessage delete message from FIFO queue
func (h *SQSHandler) DeleteMessage(result *sqs.ReceiveMessageOutput) {
	// Delete the message
	resultDelete, err := h.Service.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      h.SQSURL,
		ReceiptHandle: result.Messages[0].ReceiptHandle,
	})

	if err != nil {
		log.Println("Delete Error", err)
		return
	}

	log.Println("Message Deleted", *resultDelete)
}
