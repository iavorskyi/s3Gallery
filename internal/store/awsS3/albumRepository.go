package awsS3

import (
	"github.com/aws/aws-sdk-go/aws"
	awsS3 "github.com/aws/aws-sdk-go/service/s3"
)

type AlbumRepository struct {
	client *AWSStore
}

func (r *AlbumRepository) ListAlbums() ([]*awsS3.Bucket, error) {
	results, err := r.client.store.ListBuckets(&awsS3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	return results.Buckets, nil
}

func (r *AlbumRepository) CreateAlbum(albumName string) error {
	// For OpenStack object storage we should set an empty string to LocalizationConstraint param
	_, err := r.client.store.CreateBucket(&awsS3.CreateBucketInput{Bucket: aws.String(albumName), CreateBucketConfiguration: &awsS3.CreateBucketConfiguration{LocationConstraint: aws.String("")}})
	if err != nil {
		return err
	}
	return nil
}
