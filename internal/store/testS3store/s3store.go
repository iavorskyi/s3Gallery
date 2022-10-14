package testS3store

import (
	"github.com/iavorskyi/s3gallery/internal/model"
	"github.com/iavorskyi/s3gallery/internal/store"
)

type AWSStore struct {
	itemRepository  *ItemRepository
	albumRepository *AlbumRepository
}

func New() *AWSStore {
	return &AWSStore{}
}

func (s *AWSStore) Item() store.ItemRepository {
	if s.itemRepository != nil {
		return s.itemRepository
	}
	s.itemRepository = &ItemRepository{
		client: s,
		items:  make(map[string]*model.Item),
	}
	return s.itemRepository
}

func (s *AWSStore) Album() store.AlbumRepository {
	if s.itemRepository != nil {
		return s.albumRepository
	}
	s.itemRepository = &ItemRepository{
		client: s,
		items:  make(map[string]*model.Item),
	}
	return s.albumRepository
}
