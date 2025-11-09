package main

import (
	"context"
	"log"

	"gorm.io/driver/sqlite"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-tutorial/models"
	"github.com/FrancoLiberali/cql/logger"
)

func main() {
	// connect to db
	db, err := NewDBConnection()
	if err != nil {
		log.Fatalln(err)
	}

	// createData(db)

	// execute tutorial
	tutorial(db)

}

func NewDBConnection() (*cql.DB, error) {
	return cql.Open(
		sqlite.Open("sqlite:db"),
		&cql.Config{Logger: logger.Default.ToLogMode(logger.Info)},
	)
}

func createData(db *cql.DB) {
	ctx := context.Background()

	if err := db.GormDB.AutoMigrate(
		models.City{},
		models.Country{},
	); err != nil {
		log.Panicln(err)
	}

	usa := models.Country{
		Name: "United States of America",
	}
	if _, err := cql.Insert(ctx, db, &usa).Exec(); err != nil {
		log.Panicln(err)
	}

	parisUSA := models.City{
		CountryID:  usa.ID,
		Name:       "Paris",
		Population: 25171,
	}
	if _, err := cql.Insert(ctx, db, &parisUSA).Exec(); err != nil {
		log.Panicln(err)
	}

	washington := models.City{
		CountryID:  usa.ID,
		Name:       "Washington D. C.",
		Population: 689545,
	}
	if _, err := cql.Insert(ctx, db, &washington).Exec(); err != nil {
		log.Panicln(err)
	}

	usa.CapitalID = washington.ID
	if err := db.GormDB.Save(&usa).Error; err != nil {
		log.Panicln(err)
	}

	france := models.Country{
		Name: "France",
	}
	if _, err := cql.Insert(ctx, db, &france).Exec(); err != nil {
		log.Panicln(err)
	}

	parisFrance := models.City{
		CountryID:  france.ID,
		Name:       "Paris",
		Population: 2161000,
	}
	if _, err := cql.Insert(ctx, db, &parisFrance).Exec(); err != nil {
		log.Panicln(err)
	}

	france.CapitalID = parisFrance.ID
	if err := db.GormDB.Save(&france).Error; err != nil {
		log.Panicln(err)
	}
}
