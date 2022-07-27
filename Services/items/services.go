package items

import (
	"github.com/aws/aws-sdk-go/aws"
	awsS3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	s3Gallery "github.com/iavorskyi/s3gallery"
	"github.com/iavorskyi/s3gallery/Services/s3"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"strings"
)

func ListItems(id string) ([]s3Gallery.Item, int, error) {
	client, err := s3.GetClient()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// Get the list of items
	results, err := client.ListObjectsV2(&awsS3.ListObjectsV2Input{Bucket: aws.String(id)})
	if err != nil {
		log.Println(err)
		return nil, http.StatusInternalServerError, err
	}

	var resultList []s3Gallery.Item
	for _, i := range results.Contents {
		if strings.Contains(*i.Key, "resized/") {
			itemToList, _, err := GetItem(id, *i.Key)
			if err != nil {
				return nil, http.StatusBadRequest, err
			}
			resultList = append(resultList, itemToList)
		}

	}

	return resultList, http.StatusOK, nil
}

func GetItem(bucketID, itemID string) (s3Gallery.Item, int, error) {
	var item *awsS3.GetObjectOutput

	client, err := s3.GetClient()
	if err != nil {
		return s3Gallery.Item{}, http.StatusInternalServerError, err
	}

	// Get the list of items
	item, err = client.GetObject(&awsS3.GetObjectInput{Bucket: aws.String(bucketID), Key: aws.String(itemID)})
	if err != nil {
		log.Println(err.Error())
		return s3Gallery.Item{}, http.StatusInternalServerError, err
	}
	var itemSummary = s3Gallery.Item{Name: itemID, Album: bucketID}
	if item.ContentType != nil {
		itemSummary.Type = *item.ContentType
	}
	if item.LastModified != nil {
		itemSummary.LastModified = *item.LastModified
	}
	return itemSummary, http.StatusOK, nil
}

func UploadItem(opts UploadOpts, bucketID string) error {
	err := upload(opts, bucketID)
	if err != nil {
		return err
	}

	tempFile, err := os.Create("./rzdImg" + opts.Name + ".jpeg")
	if err != nil {
		return err
	}
	defer os.Remove("./rzdImg" + opts.Name + ".jpeg")

	fileToResize, err := os.Open(opts.Path)
	if err != nil {
		return err
	}
	defer fileToResize.Close()

	img, _, err := image.Decode(fileToResize)
	var newImg image.Image
	newImg = resize.Resize(0, 280, img, resize.Lanczos3)

	err = jpeg.Encode(tempFile, newImg, nil)
	err = tempFile.Close()
	if err != nil {
		return err
	}

	resizedImgOpts := UploadOpts{Name: "resized/" + opts.Name, Path: "./rzdImg" + opts.Name + ".jpeg"}
	err = upload(resizedImgOpts, bucketID)
	if err != nil {
		return err
	}

	return err
}

func DeleteItem(bucket, item string) error {
	client, err := s3.GetClient()
	if err != nil {
		return err
	}
	_, err = client.DeleteObject(&awsS3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(item),
	})
	if err != nil {
		return err
	}

	// delete resized copy
	_, err = client.DeleteObject(&awsS3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("resized/" + item),
	})
	if err != nil {
		return err
	}

	// Wait to assure item is deleted
	err = client.WaitUntilObjectNotExists(&awsS3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(item),
	})
	if err != nil {
		return err
	}
	// Wait to assure resized item is deleted
	err = client.WaitUntilObjectNotExists(&awsS3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("resized/" + item),
	})
	if err != nil {
		return err
	}

	return nil
}

func upload(opts UploadOpts, bucket string) error {
	file, err := os.Open(opts.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	uploader, err := s3.GetManager()
	item := &s3manager.UploadInput{
		Bucket:   aws.String(bucket),
		Key:      aws.String(opts.Name),
		Body:     file,
		Metadata: map[string]*string{"format": aws.String(opts.Format)},
		ACL:      aws.String("public-read"),
	}
	_, err = uploader.Upload(item)
	return err
}
