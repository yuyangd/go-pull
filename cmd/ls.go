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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/spf13/cobra"
	"github.com/yuyangd/go-pull/receiver"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "Listing objects in the source bucket",
	Example: `go-pull ls`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Listing objects in the source bucket")
		s3 := receiver.S3Client()
		(&receiver.S3Handler{
			BucketName: aws.String(BucketName),
			Service:    s3,
		}).ListObjects()
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
