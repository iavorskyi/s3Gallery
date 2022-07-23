package items

import "github.com/go-pg/pg/v10"

type ItemClient struct {
	DB *pg.DB
}

type Item struct {
	ID string `json:"id"`
}

type UploadOpts struct {
	Name   string
	Title  string
	Path   string
	Format string
}
