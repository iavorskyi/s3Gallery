package sqlStore

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"testing"
)

// TestStore ...
func TestDB(t *testing.T, connString, dbUser, dbName, dbPassword string) (*pg.DB, func(...string)) {
	t.Helper()
	db := pg.Connect(&pg.Options{
		Addr:     connString,
		User:     dbUser,
		Password: dbPassword,
		Database: dbName,
	})
	if err := db.Ping(context.Background()); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			for _, table := range tables {
				query := fmt.Sprintf("TRUNCATE TABLE %s", table)
				if _, err := db.Exec(query); err != nil {
					t.Fatal(err)
				}
			}
		}
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}

}
