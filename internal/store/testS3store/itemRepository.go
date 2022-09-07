package testS3store

import (
	"errors"
	"github.com/iavorskyi/s3gallery/internal/model"
)

type ItemRepository struct {
	client *AWSStore
	items  map[string]*model.Item
}

func (r *ItemRepository) ListItemsByAlbumId(id string) ([]model.Item, error) {
	var results []model.Item
	if r.items == nil {
		return results, nil
	}
	for _, itm := range r.items {
		if itm.Album == id {
			results = append(results, *itm)
		}
	}
	return results, nil
}

func (r *ItemRepository) GetItemByAlbumIDAndID(albumID, id string) (model.Item, error) {
	if len(r.items) == 0 {
		return model.Item{}, errors.New("no item with this name")
	}
	for _, itm := range r.items {
		if itm.Album == albumID {
			if itm.Name == id {
				return *itm, nil
			}
		}
	}
	return model.Item{}, errors.New("no item with this name")
}
func (r *ItemRepository) UploadItem(fileName string, path string, bucket string) (string, error) {
	item := model.Item{Name: fileName, Album: bucket}
	r.items[item.Name] = &item
	return item.Name, nil
}

func (r *ItemRepository) DeleteItemByBucketIDAndItemID(itemID, bucket string) error {
	for _, itm := range r.items {
		if itm.Album == bucket {
			if itm.Name == itemID {
				delete(r.items, itm.Name)
			}
		}
	}
	return nil
}
