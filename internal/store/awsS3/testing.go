package awsS3

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"testing"
)

//TODO CREATE TEST S3 STORE. THROW

// TestAWSStoreClient ...
func TestAWSStoreClient(t *testing.T) (*s3.S3, error) {
	return GetClient()
}

// TestAWSStoreClient ...
func TestAWSStoreManager(t *testing.T) (*s3manager.Uploader, error) {
	return GetManager()
}
