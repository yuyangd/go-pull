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

package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/yuyangd/go-pull/receiver"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "get the next object in the queue",
	Example: `go-pull get`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Get the next object in the queue")
		sqs := receiver.SqsClient()
		s3 := receiver.S3Client()
		qh := &receiver.SQSHandler{
			SQSURL:  &SQSURL,
			Service: sqs,
		}
		result, err := qh.ReceiveMessage()

		if err != nil {
			log.Println("Error receiving the message")
		}
		// Receive bucket and key from the Queue message
		keyP := result.Messages[0].MessageAttributes["Key"].StringValue
		bucketP := result.Messages[0].MessageAttributes["Bucket"].StringValue

		// Delete the message
		qh.DeleteMessage(result)

		// Download the object
		// if download failed due to object not exists
		// -> re-request another message from the Queue
		sh := &receiver.S3Handler{
			BucketName: bucketP,
			Service:    s3,
		}
		sh.GetObject(keyP)

		// Delete the object
		sh.DeleteObject(keyP)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
