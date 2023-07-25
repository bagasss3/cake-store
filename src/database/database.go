package database

import (
	log "github.com/sirupsen/logrus"

	"cake-store/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", config.DBDSN())
	if err != nil {
		log.WithField("dbDSN", config.DBDSN()).Fatal("Failed to connect:", err)
	}

	db.SetMaxIdleConns(config.MaxIdleConns())
	db.SetMaxOpenConns(config.MaxOpenConns())
	db.SetConnMaxLifetime(config.ConnMaxLifeTime())
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime())

	log.Info("Success connect database")
	return db
}
