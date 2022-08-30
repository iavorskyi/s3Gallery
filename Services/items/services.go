package items

import (
	"github.com/aws/aws-sdk-go/aws"
	awsS3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/iavorskyi/s3gallery/internal/store"
	awsStore "github.com/iavorskyi/s3gallery/internal/store/awsS3"
	"github.com/iavorskyi/s3gallery/s3Gallery"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"strings"
)

func ListItems(id string, store store.S3Store) ([]s3Gallery.Item, int, error) {
	// Get the list of items
	results, err := store.Item().ListItemsByAlbumId(id)
	if err != nil {
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

	client, err := awsStore.GetClient()
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

func UploadItem(fileName, path, bucketID string) error {
	err := upload(fileName, path, bucketID)
	if err != nil {
		return err
	}

	tempFile, err := os.Create("./rzdImg" + fileName + ".jpeg")
	if err != nil {
		return err
	}
	defer os.Remove("./rzdImg" + fileName + ".jpeg")

	fileToResize, err := os.Open(path)
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

	resizedFileName := "resized/" + fileName
	resizedFilePath := "./rzdImg" + fileName + ".jpeg"
	err = upload(resizedFileName, resizedFilePath, bucketID)
	if err != nil {
		return err
	}

	return err
}

func DeleteItem(bucket, item string) error {
	client, err := awsStore.GetClient()
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

func upload(fileName, path, bucket string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	uploader, err := awsStore.GetManager()
	item := &s3manager.UploadInput{
		Bucket:   aws.String(bucket),
		Key:      aws.String(fileName),
		Body:     file,
		Metadata: map[string]*string{"format": aws.String(".jpeg")},
		ACL:      aws.String("public-read"),
	}
	_, err = uploader.Upload(item)
	return err
}
