//go:build tutorial_8
// +build tutorial_8

package main

import (
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-tutorial/conditions"
	"github.com/FrancoLiberali/cql-tutorial/models"
	"gorm.io/gorm"
)

// Target: create Rennes in France and then delete it
func tutorial(db *gorm.DB) {
	// get country called France
	france, err := cql.Query[models.Country](
		db,
		conditions.Country.Name.Is().Eq("France"),
	).FindOne()

	if err != nil {
		log.Panicln(err)
	}

	// create Rennes using gorm's Create method
	rennes := models.City{
		Country:    france,
		Name:       "Rennes",
		Population: 215366,
	}
	if err := db.Create(&rennes).Error; err != nil {
		log.Panicln(err)
	}

	// delete city called Rennes
	deleted, err := cql.Delete[models.City](
		db,
		conditions.City.Name.Is().Eq("Rennes"),
	).Exec()

	// SQL executed:
	// UPDATE cities
	// SET deleted_at="2023-09-11 10:54:12.598"
	// WHERE cities.name = "Rennes" AND cities.deleted_at IS NULL

	if err != nil {
		log.Panicln(err)
	}

	log.Printf("Deleted %v city\n", deleted)
}
