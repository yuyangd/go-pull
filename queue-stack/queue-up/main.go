package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	qURL = os.Getenv("SQS_URL")
)

func getService() *sqs.SQS {
	// Create Session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create a SQS client from just a session.
	return sqs.New(sess)
}

func sendMessage(svc *sqs.SQS, key *string, bucket *string, t time.Time) {
	result, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageGroupId: aws.String("models"),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Bucket": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: bucket,
			},
			"Key": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: key,
			},
		},
		MessageBody: aws.String(t.String()),
		QueueUrl:    &qURL,
	})
	if err != nil {
		log.Println("Error", err)
	}
	log.Println("Success", *result.MessageId)
}

func handler(ctx context.Context, s3Event events.S3Event) {
	for _, record := range s3Event.Records {
		s3 := record.S3
		log.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
		sendMessage(getService(), aws.String(s3.Object.Key), aws.String(s3.Bucket.Name), record.EventTime)
	}
}

func main() {
	lambda.Start(handler)
}
