package sqlStore_test

import (
	"os"
	"testing"
)

var (
	dbConnectionString string
	dbName             string
	dbUser             string
	dbPassword         string
)

func TestMain(m *testing.M) {
	dbConnectionString = os.Getenv("TEST_DB_CONN_STR")
	if dbConnectionString == "" {
		dbConnectionString = "185.226.41.166:5432"
	}
	dbName = os.Getenv("TEST_DB_NAME")
	if dbName == "" {
		dbName = "S3GalleryTest"
	}
	dbUser = os.Getenv("TEST_DB_USER")
	if dbUser == "" {
		dbUser = "admin"
	}
	dbPassword = os.Getenv("TEST_DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "QDiW2auJIOxv7Zm9C8kz643GUMgSh051"
	}
	os.Exit(m.Run())
}
