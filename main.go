package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql/logger"
)

func main() {
	// connect to db
	db, err := NewDBConnection()
	if err != nil {
		log.Fatalln(err)
	}

	// execute tutorial
	tutorial(db)
}

func NewDBConnection() (*gorm.DB, error) {
	return cql.Open(
		sqlite.Open("sqlite:db"),
		&gorm.Config{Logger: logger.Default.ToLogMode(logger.Info)},
	)
}
