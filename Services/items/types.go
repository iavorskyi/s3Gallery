package items

import (
	"github.com/go-pg/pg/v10"
)

type ItemClient struct {
	DB *pg.DB
}

type Item struct {
	ID string `json:"id"`
}
