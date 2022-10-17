package items

import (
	"github.com/iavorskyi/s3gallery/internal/model"
	"github.com/iavorskyi/s3gallery/internal/store"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"net/http"
	"os"
)

func ListItems(id string, s3store store.S3Store) ([]model.Item, int, error) {
	// Get the list of items
	items, err := s3store.Item().ListItemsByAlbumId(id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return items, http.StatusOK, nil
}

func GetItem(bucketID, itemID string, s3Store store.S3Store) (model.Item, int, error) {
	item, err := s3Store.Item().GetItemByAlbumIDAndID(bucketID, itemID)
	if err != nil {
		return item, http.StatusInternalServerError, err
	}
	return item, http.StatusOK, nil
}

func UploadItem(fileName, path, bucketID string, s3Store store.S3Store) error {
	_, err := s3Store.Item().UploadItem(fileName, path, bucketID)
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
	_, err = s3Store.Item().UploadItem(resizedFileName, resizedFilePath, bucketID)
	if err != nil {
		return err
	}

	return err
}

func DeleteItem(bucket, item string, s3store store.S3Store) error {
	err := s3store.Item().DeleteItemByBucketIDAndItemID(bucket, item)
	if err != nil {
		return err
	}
	err = s3store.Item().DeleteItemByBucketIDAndItemID(bucket, "resized/"+item)
	if err != nil {
		return err
	}
	return nil
}
