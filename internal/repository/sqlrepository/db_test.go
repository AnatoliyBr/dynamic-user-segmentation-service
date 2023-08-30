package sqlrepository_test

import (
	"os"
	"testing"
)

var (
	testDatabaseURL string
)

func TestMain(m *testing.M) {
	testDatabaseURL = os.Getenv("TEST_DATABASE_URL")
	if testDatabaseURL == "" {
		testDatabaseURL = "host=localhost user=dev password=qwerty dbname=user_seg_app_test sslmode=disable"
	}

	os.Exit(m.Run())
}
