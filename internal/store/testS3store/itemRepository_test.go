package testS3store_test

import (
	"github.com/iavorskyi/s3gallery/internal/model"
	"github.com/iavorskyi/s3gallery/internal/store/testS3store"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestItemRepository_ListItemsByAlbumId(t *testing.T) {
	s := testS3store.New()
	item := model.TestItem(t)
	s.Item().UploadItem("test", "test", "test")

	items, err := s.Item().ListItemsByAlbumId(item.Album)
	assert.NoError(t, err)
	log.Println(items)
	assert.NotNil(t, items)
}
