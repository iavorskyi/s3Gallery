package awsS3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/iavorskyi/s3gallery/internal/store"
)

import "os"

type AWSStore struct {
	store          *s3.S3
	manager        *s3manager.Uploader
	itemRepository *ItemRepository
}

func New(s3 *s3.S3, manager *s3manager.Uploader) *AWSStore {
	return &AWSStore{store: s3, manager: manager}
}

// DirectoryIterator represents an iterator of a specified directory
type DirectoryIterator struct {
	filePaths []string
	bucket    string
	next      struct {
		path string
		f    *os.File
	}
	err error
}

func GetClient() (*s3.S3, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	sess, _ = session.NewSession(&aws.Config{
		Region:           aws.String("us-east-2"),
		Endpoint:         aws.String("https://upper-austria.ventuscloud.eu:8080"),
		S3ForcePathStyle: aws.Bool(true)},
	)

	svc := s3.New(sess)
	return svc, nil
}

func GetManager() (*s3manager.Uploader, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	sess, _ = session.NewSession(&aws.Config{
		Region:           aws.String("any"),
		Endpoint:         aws.String("https://upper-austria.ventuscloud.eu:8080"),
		S3ForcePathStyle: aws.Bool(true)},
	)

	service := s3manager.NewUploader(sess)
	return service, nil
}

func (s *AWSStore) Item() store.ItemRepository {
	if s.itemRepository != nil {
		return s.itemRepository
	}
	s.itemRepository = &ItemRepository{
		client: s,
	}
	return s.itemRepository
}
