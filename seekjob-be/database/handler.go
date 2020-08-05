package database

import (
	"fmt"
	"log"
	"seekjob/config"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var singletonHandler *gorm.DB

func init() {
	postgresCfg := config.Config.PostgresCfg
	postgresString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		postgresCfg.Host,
		postgresCfg.User,
		postgresCfg.Password,
		postgresCfg.DbName,
		postgresCfg.SSLMode,
	)
	db, err := gorm.Open("postgres", postgresString)
	if err != nil {
		log.Fatalf("[ERROR] Fatal error connecting database: %s", err)
		return
	}

	db.LogMode(true)
	singletonHandler = db
}

// GetHandler returns the handler of the database
func GetHandler() *gorm.DB {
	return singletonHandler
}
