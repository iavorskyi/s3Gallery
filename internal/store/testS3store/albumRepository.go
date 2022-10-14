package testS3store

import (
	"github.com/aws/aws-sdk-go/aws"
	awsS3 "github.com/aws/aws-sdk-go/service/s3"
	"time"
)

type AlbumRepository struct {
	client *AWSStore
	albums map[string]*awsS3.Bucket
}

func (r *AlbumRepository) ListAlbums() ([]*awsS3.Bucket, error) {
	albums := []*awsS3.Bucket{}
	for _, v := range r.albums {
		albums = append(albums, v)
	}
	return albums, nil
}

func (r *AlbumRepository) CreateAlbum(albumName string) error {
	r.albums[albumName] = &awsS3.Bucket{
		Name:         aws.String(albumName),
		CreationDate: aws.Time(time.Now()),
	}
	return nil
}
