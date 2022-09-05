package store

import (
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
