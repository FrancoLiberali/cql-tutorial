package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ditrit/badaas/orm"
	"github.com/ditrit/badaas/orm/logger"
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
	return orm.Open(
		sqlite.Open(orm.CreateSQLiteDSN("db")),
		&gorm.Config{Logger: logger.Default.ToLogMode(logger.Info)},
	)
}
