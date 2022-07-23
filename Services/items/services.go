package items

import (
	"github.com/aws/aws-sdk-go/aws"
	awsS3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/iavorskyi/s3gallery/Services/s3"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
)

func ListItems(id string) ([]*awsS3.GetObjectOutput, int, error) {
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

	var resultList = []*awsS3.GetObjectOutput{}
	for _, i := range results.Contents {
		itemToList, _, err := GetItem(id, *i.Key)
		if err != nil {
			return nil, http.StatusBadRequest, err
		}
		resultList = append(resultList, itemToList)
	}

	return resultList, http.StatusOK, nil
}

func GetItem(bucketID, itemID string) (*awsS3.GetObjectOutput, int, error) {
	var item *awsS3.GetObjectOutput
	client, err := s3.GetClient()
	if err != nil {
		return item, http.StatusInternalServerError, err
	}

	// Get the list of items
	item, err = client.GetObject(&awsS3.GetObjectInput{Bucket: aws.String(bucketID), Key: aws.String(itemID)})
	if err != nil {
		log.Println(err.Error())
		return item, http.StatusInternalServerError, err
	}
	return item, http.StatusOK, nil
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
	newImg = resize.Resize(160, 160, img, resize.Lanczos3)

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

	uploader2, err := s3.GetManager()
	resizedItem := &s3manager.UploadInput{
		Bucket:   aws.String(bucket),
		Key:      aws.String(opts.Name),
		Body:     file,
		Metadata: map[string]*string{"format": aws.String(opts.Format)},
	}
	_, err = uploader2.Upload(resizedItem)
	return err
}
