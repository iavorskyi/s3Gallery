package store

type S3Store interface {
	Item() ItemRepository
}
