// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	// Create Session
	// sess := session.Must(session.NewSession())
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create a SQS client from just a session.
	svc := sqs.New(sess)

	qURL := "https://sqs.ap-southeast-2.amazonaws.com/567418462583/test-fifo-queue-ModelUpdatesSQSQueue-13LVZGQ6PGZTS.fifo"

	result, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageGroupId: aws.String("models"),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Bucket": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("Racing-Model1"),
			},
			"Key": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("racer-v6"),
			},
		},
		MessageDeduplicationId: aws.String("abcd3"),
		MessageBody:            aws.String("racer-v6"),
		QueueUrl:               &qURL,
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success", *result.MessageId)

}
