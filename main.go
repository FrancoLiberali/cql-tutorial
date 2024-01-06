package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

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

func NewDBConnection() (*gorm.DB, error) {
	return cql.Open(
		sqlite.Open("sqlite:db"),
		&gorm.Config{Logger: logger.Default.ToLogMode(logger.Info)},
	)
}

func createData(db *gorm.DB) {
	if err := db.AutoMigrate(
		models.City{},
		models.Country{},
	); err != nil {
		log.Panicln(err)
	}

	usa := models.Country{
		Name: "United States of America",
	}
	if err := db.Create(&usa).Error; err != nil {
		log.Panicln(err)
	}

	parisUSA := models.City{
		Country:    &usa,
		Name:       "Paris",
		Population: 25171,
	}
	if err := db.Create(&parisUSA).Error; err != nil {
		log.Panicln(err)
	}

	washington := models.City{
		Country:    &usa,
		Name:       "Washington D. C.",
		Population: 689545,
	}
	if err := db.Create(&washington).Error; err != nil {
		log.Panicln(err)
	}

	usa.CapitalID = washington.ID
	if err := db.Save(&usa).Error; err != nil {
		log.Panicln(err)
	}

	france := models.Country{
		Name: "France",
	}
	if err := db.Create(&france).Error; err != nil {
		log.Panicln(err)
	}

	parisFrance := models.City{
		Country:    &france,
		Name:       "Paris",
		Population: 2161000,
	}
	if err := db.Create(&parisFrance).Error; err != nil {
		log.Panicln(err)
	}

	france.CapitalID = parisFrance.ID
	if err := db.Save(&france).Error; err != nil {
		log.Panicln(err)
	}
}
