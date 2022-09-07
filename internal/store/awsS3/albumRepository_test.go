package awsS3_test

import (
	"github.com/iavorskyi/s3gallery/internal/store/awsS3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlbumRepository_ListAlbums(t *testing.T) {
	client, err := awsS3.TestAWSStoreClient(t)
	if err != nil {
		t.Fatal(err)
	}

	s3Store := awsS3.New(client, nil)
	albums, err := s3Store.Album().ListAlbums()
	assert.NoError(t, err)
	assert.NotNil(t, albums)
}

func TestAlbumRepository_CreateAlbum(t *testing.T) {
	client, err := awsS3.TestAWSStoreClient(t)
	if err != nil {
		t.Fatal(err)
	}

	//TODO remove created album after testing
	s3Store := awsS3.New(client, nil)
	err = s3Store.Album().CreateAlbum("test")
	assert.NoError(t, err)
}
