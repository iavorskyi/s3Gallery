package awsS3

import (
	"github.com/aws/aws-sdk-go/aws"
	awsS3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/iavorskyi/s3gallery/internal/model"
	"log"
	"os"
	"strings"
)

type ItemRepository struct {
	client *AWSStore
}

func (r *ItemRepository) ListItemsByAlbumId(id string) ([]model.Item, error) {
	results, err := r.client.store.ListObjectsV2(&awsS3.ListObjectsV2Input{Bucket: aws.String(id)})
	if err != nil {
		return nil, err
	}

	var resultList []model.Item
	for _, i := range results.Contents {
		if strings.Contains(*i.Key, "resized/") {
			itemToList, err := r.GetItemByAlbumIDAndID(id, *i.Key)
			if err != nil {
				return nil, err
			}
			resultList = append(resultList, itemToList)
		}
	}
	return resultList, nil
}

func (r *ItemRepository) GetItemByAlbumIDAndID(bucketID, id string) (model.Item, error) {
	item, err := r.client.store.GetObject(&awsS3.GetObjectInput{Bucket: aws.String(bucketID), Key: aws.String(id)})
	if err != nil {
		log.Println(err.Error())
		return model.Item{}, err
	}
	var itemSummary = model.Item{Name: id, Album: bucketID}
	if item.ContentType != nil {
		itemSummary.Type = *item.ContentType
	}
	if item.LastModified != nil {
		itemSummary.LastModified = *item.LastModified
	}
	return itemSummary, nil
}

func (r *ItemRepository) UploadItem(fileName, path, bucket string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	item := &s3manager.UploadInput{
		Bucket:   aws.String(bucket),
		Key:      aws.String(fileName),
		Body:     file,
		Metadata: map[string]*string{"format": aws.String(".jpeg")},
		ACL:      aws.String("public-read"),
	}
	_, err = r.client.manager.Upload(item)
	return fileName, err
}

func (r *ItemRepository) DeleteItemByBucketIDAndItemID(bucketID, itemID string) error {
	_, err := r.client.store.DeleteObject(&awsS3.DeleteObjectInput{
		Bucket: aws.String(bucketID),
		Key:    aws.String(itemID),
	})
	if err != nil {
		return err
	}

	err = r.client.store.WaitUntilObjectNotExists(&awsS3.HeadObjectInput{
		Bucket: aws.String(bucketID),
		Key:    aws.String(itemID),
	})
	return nil
}
