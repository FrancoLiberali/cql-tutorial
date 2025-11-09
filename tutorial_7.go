//go:build tutorial_7
// +build tutorial_7

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-tutorial/conditions"
	"github.com/FrancoLiberali/cql-tutorial/models"
)

// Target: edit Paris France population to 2102650
func tutorial(db *cql.DB) {
	var cities []models.City

	updated, err := cql.Update[models.City](
		context.Background(),
		db,
		conditions.City.Name.Is().Eq(cql.String("Paris")),
		conditions.City.Country(
			conditions.Country.Name.Is().Eq(cql.String("France")),
		),
	).Returning(&cities).Set(
		conditions.City.Population.Set().Eq(cql.Int64(2102650)),
	)

	// SQL executed:
	// UPDATE cities
	// SET population=2102650
	// FROM countries Country
	// WHERE cities.name = "Paris" AND
	//    (Country.id = cities.country_id AND Country.name = "France")
	// RETURNING *

	if err != nil {
		log.Panicln(err)
	}

	parisFrance := cities[0]
	fmt.Println("--------------------------")
	fmt.Printf("Updated %v city: %v\n", updated, parisFrance)
	fmt.Println("Initial population was 2161000")

	// go back to initial situation with gorm's Save method
	parisFrance.Population = 2161000
	if err := db.GormDB.Save(&parisFrance).Error; err != nil {
		log.Panicln(err)
	}
}
