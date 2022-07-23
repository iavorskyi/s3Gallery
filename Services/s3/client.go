package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

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
