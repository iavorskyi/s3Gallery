package awsS3_test

import (
	"github.com/iavorskyi/s3gallery/internal/store/awsS3"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestItemRepository_ListItemsByAlbumId(t *testing.T) {
	client, err := awsS3.TestAWSStoreClient(t)
	if err != nil {
		t.Fatal(err)
	}

	s3Store := awsS3.New(client, nil)
	_, err = s3Store.Item().ListItemsByAlbumId("wrong")
	assert.Error(t, err)
	item, err := s3Store.Item().ListItemsByAlbumId("demo")
	assert.NoError(t, err)
	assert.NotNil(t, item)
}

func TestItemRepository_GetItemByAlbumIDAndID(t *testing.T) {
	client, err := awsS3.TestAWSStoreClient(t)
	if err != nil {
		t.Fatal(err)
	}

	// Wrong item ID
	s3Store := awsS3.New(client, nil)
	_, err = s3Store.Item().GetItemByAlbumIDAndID("demo", "test")
	assert.Error(t, err)
}

func TestItemRepository_UploadItem(t *testing.T) {
	manger, err := awsS3.TestAWSStoreManager(t)
	if err != nil {
		t.Fatal(err)
	}
	// Create empty image file
	fileName := "image.jpeg"
	os.Create(fileName)
	defer os.Remove(fileName)

	s3Store := awsS3.New(nil, manger)
	uploadedFileName, err := s3Store.Item().UploadItem(fileName, fileName, "demo")
	assert.NoError(t, err)
	assert.NotEmpty(t, uploadedFileName)

}

func TestItemRepository_DeleteItem(t *testing.T) {
	client, err := awsS3.TestAWSStoreClient(t)
	if err != nil {
		t.Fatal(err)
	}
	manger, err := awsS3.TestAWSStoreManager(t)
	if err != nil {
		t.Fatal(err)
	}

	// Create empty image file
	fileName := "image.jpeg"
	os.Create(fileName)
	defer os.Remove(fileName)

	s3Store := awsS3.New(client, manger)

	_, err = s3Store.Item().UploadItem(fileName, fileName, "demo")
	if err != nil {
		t.Fatal(err)
	}

	err = s3Store.Item().DeleteItemByBucketIDAndItemID("demo", fileName)
	assert.NoError(t, err)

}
