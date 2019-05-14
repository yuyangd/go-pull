package receiver

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//S3Handler data for s3 service
type S3Handler struct {
	Service    S3Iface
	BucketName *string
}

// S3Iface defines AWS s3 APIs
type S3Iface interface {
	ListObjects(*s3.ListObjectsInput) (*s3.ListObjectsOutput, error)
}

// S3Client creates a client from session
func S3Client() *s3.S3 {
	// Create Session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create a S3 client from this session.
	return s3.New(sess)
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
