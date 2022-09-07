package albums

import (
	"errors"
	awsS3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/iavorskyi/s3gallery/internal/store"
	"net/http"
)

func ListAlbums(s3store store.S3Store) ([]*awsS3.Bucket, int, error) {
	albums, err := s3store.Album().ListAlbums()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return albums, http.StatusOK, nil
}

func GetAlbum(albumID string, s3store store.S3Store) (*awsS3.Bucket, int, error) {
	albums, err := s3store.Album().ListAlbums()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var bucket awsS3.Bucket
	for _, b := range albums {
		if *b.Name == albumID {
			bucket = *b
		}
	}
	if bucket.Name == nil {
		return nil, http.StatusNotFound, errors.New("bucket is not found")
	}
	return &bucket, http.StatusOK, nil
}

func CreateAlbum(albumName string, s3store store.S3Store) (int, error) {

	err := s3store.Album().CreateAlbum(albumName)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
