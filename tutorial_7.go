//go:build tutorial_7
// +build tutorial_7

package main

import (
	"log"

	"github.com/ditrit/badaas-orm-tutorial/conditions"
	"github.com/ditrit/badaas-orm-tutorial/models"
	"github.com/ditrit/badaas/orm"
	"gorm.io/gorm"
)

// Target: edit Paris France population to 2102650
func tutorial(db *gorm.DB) {
	var cities []models.City

	updated, err := orm.Update[models.City](
		db,
		conditions.City.Name.Is().Eq("Paris"),
		conditions.City.Country(
			conditions.Country.Name.Is().Eq("France"),
		),
	).Returning(&cities).Set(
		conditions.City.Population.Set().Eq(2102650),
	)

	// SQL executed:
	// UPDATE cities
	// SET population=2102650,updated_at="2023-09-11 10:41:19.272"
	// FROM countries Country
	// WHERE cities.name = "Paris" AND
	//    (Country.id = cities.country_id AND Country.name = "France" AND Country.deleted_at IS NULL) AND
	//    cities.deleted_at IS NULL
	// RETURNING *

	if err != nil {
		log.Panicln(err)
	}

	parisFrance := cities[0]
	log.Printf("Updated %v city: %v\n", updated, parisFrance)
	log.Println("Initial population was 2161000")

	// go back to initial situation with gorm's Save method
	parisFrance.Population = 2161000
	db.Save(&parisFrance)
}
