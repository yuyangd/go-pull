package receiver

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//S3Handler data for s3 service
type S3Handler struct {
	Service    S3Iface
	BucketName *string
}

// S3Iface defines AWS s3 APIs
type S3Iface interface {
	ListObjects(*s3.ListObjectsInput) (*s3.ListObjectsOutput, error)
	DeleteObject(*s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error)
	GetObject(*s3.GetObjectInput) (*s3.GetObjectOutput, error)
}

// S3Client creates a client from session
func S3Client() *s3.S3 {
	// Create Session
	sess := s3Session()

	// Create a S3 client from this session.
	return s3.New(sess)
}

func s3Session() *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
}

// ListObjects equals to aws s3 ls
func (h *S3Handler) ListObjects() {
	params := &s3.ListObjectsInput{
		Bucket: h.BucketName,
	}
	resp, _ := h.Service.ListObjects(params)
	for _, key := range resp.Contents {
		fmt.Println(*key.Key)
	}
}

// DeleteObject from non-versioned bucket
func (h *S3Handler) DeleteObject(key *string) (err error) {
	input := &s3.DeleteObjectInput{
		Bucket: h.BucketName,
		Key:    key,
	}
	_, err = h.Service.DeleteObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Println(aerr.Error())
			}
		} else {
			log.Println(err.Error())
		}
		return
	}
	return nil
}

// GetObject downloads the object
func (h *S3Handler) GetObject(key *string) (err error) {

	downloader := s3manager.NewDownloader(s3Session())
	// Create a file to write the S3 Object contents to.
	f, err := os.Create(*key)
	if err != nil {
		return fmt.Errorf("failed to create file %q, %v", *key, err)
	}

	// Write the contents of S3 Object to the file
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: h.BucketName,
		Key:    key,
	})
	if err != nil {
		return fmt.Errorf("failed to download file, %v", err)
	}
	log.Printf("file downloaded, %d bytes\n", n)
	return nil
}
