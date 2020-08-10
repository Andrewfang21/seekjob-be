package models_test

import (
	"database/sql"
	"fmt"
	"seekjob/database"
	"strings"
	"testing"
)

func setUpDatabase(t *testing.T, tablesToTruncate ...string) *sql.DB {
	// ormer := database.GetHandler()
	ormer := database.SingletonHandler

	truncateQuery := fmt.Sprintf("TRUNCATE TABLE %s CASCADE;", strings.Join(tablesToTruncate, ", "))
	_, err := ormer.Exec(truncateQuery)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	return ormer
}
