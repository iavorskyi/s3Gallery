package albums

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	awsS3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/iavorskyi/s3gallery/Services/s3"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ListAlbums() ([]*awsS3.Bucket, int, error) {
	client, err := s3.GetClient()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	result, err := client.ListBuckets(&awsS3.ListBucketsInput{})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return result.Buckets, http.StatusOK, nil
}

func GetAlbum(albumID string) (*awsS3.Bucket, int, error) {
	client, err := s3.GetClient()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	result, err := client.ListBuckets(&awsS3.ListBucketsInput{})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var bucket awsS3.Bucket
	for _, b := range result.Buckets {
		if *b.Name == albumID {
			bucket = *b
		}
	}
	if bucket.Name == nil {
		return nil, http.StatusNotFound, errors.New("bucket is not found")
	}
	return &bucket, http.StatusOK, nil
}

func CreateAlbum(bucketName string) (int, error) {
	client, err := s3.GetClient()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	// For OpenStack object storage we should set an empty string to LocalizationConstraint param
	result, err := client.CreateBucket(&awsS3.CreateBucketInput{Bucket: aws.String(bucketName), CreateBucketConfiguration: &awsS3.CreateBucketConfiguration{LocationConstraint: aws.String("")}})
	if err != nil {
		return http.StatusInternalServerError, err
	}
	logrus.Print(result.Location)
	return http.StatusOK, nil
}
