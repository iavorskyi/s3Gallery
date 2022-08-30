package store

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/iavorskyi/s3gallery/internal/model"
)

// UserRepository ...
type UserRepository interface {
	Create(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
	FindByCredential(email, password string) (int, error)
}

type ItemRepository interface {
	ListItemsByAlbumId(id string) (*s3.ListObjectsV2Output, error)
}
