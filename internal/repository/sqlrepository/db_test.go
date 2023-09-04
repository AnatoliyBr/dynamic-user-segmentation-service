package sqlrepository_test

import (
	"os"
	"testing"
)

var (
	testDatabaseURL string
)

func TestMain(m *testing.M) {
	testDatabaseURL = os.Getenv("TEST_DATABASE_URL_LOCALHOST")
	if testDatabaseURL == "" {
		testDatabaseURL = "postgres://dev:qwerty@localhost:5432/user_seg_app_test"
	}

	testDatabaseURL += "?sslmode=disable"

	os.Exit(m.Run())
}
