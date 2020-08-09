package database

import (
	"database/sql"
	"fmt"
	"log"
	"seekjob/config"

	_ "github.com/lib/pq"
)

var singletonHandler *sql.DB

func init() {
	postgresCfg := config.Config.PostgresCfg
	postgresString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		postgresCfg.Host,
		postgresCfg.User,
		postgresCfg.Password,
		postgresCfg.DbName,
		postgresCfg.SSLMode,
	)
	db, err := sql.Open("postgres", postgresString)
	if err != nil {
		log.Fatalf("[ERROR] Fatal error connecting database: %s", err)
		return
	}

	singletonHandler = db
}

// GetHandler returns the handler of the database
func GetHandler() *sql.DB {
	return singletonHandler
}
