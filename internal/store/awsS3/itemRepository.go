package awsS3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	awsS3 "github.com/aws/aws-sdk-go/service/s3"
	"log"
)

type ItemRepository struct {
	client *AWSStore
}

func (r *ItemRepository) ListItemsByAlbumId(id string) (*s3.ListObjectsV2Output, error) {
	results, err := r.client.store.ListObjectsV2(&awsS3.ListObjectsV2Input{Bucket: aws.String(id)})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return results, nil
}
