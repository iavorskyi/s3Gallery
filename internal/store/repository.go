package store

import (
	awsS3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/iavorskyi/s3gallery/internal/model"
)

// UserRepository ...
type UserRepository interface {
	Create(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
	FindByCredential(email, password string) (int, error)
}

type ItemRepository interface {
	ListItemsByAlbumId(id string) ([]model.Item, error)
	GetItemByAlbumIDAndID(bucketID, id string) (model.Item, error)
	UploadItem(fileName, path, bucket string) (string, error)
	DeleteItemByBucketIDAndItemID(bucketID, itemID string) error
}

type AlbumRepository interface {
	ListAlbums() ([]*awsS3.Bucket, error)
	CreateAlbum(albumName string) error
}
