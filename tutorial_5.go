//go:build tutorial_5
// +build tutorial_5

package main

import (
	"fmt"
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-tutorial/conditions"
	"github.com/FrancoLiberali/cql-tutorial/models"
	"gorm.io/gorm"
)

// Target: get all cities whose name is 'Paris' and preload its country
func tutorial(db *gorm.DB) {
	cities, err := cql.Query[models.City](
		db,
		conditions.City.Name.Is().Eq("Paris"),
		conditions.City.PreloadCountry(),
	).Find()

	// Equivalent to:
	// cities, err := cql.Query[models.City](
	// 	db,
	// 	conditions.City.Name.Is().Eq("Paris"),
	// 	conditions.City.Country(
	// 		conditions.Country.Preload(),
	// 	),
	// ).Find()

	// SQL executed:
	// SELECT cities.*,
	//    Country.id AS Country__id,Country.created_at AS Country__created_at,Country.updated_at AS Country__updated_at,Country.deleted_at AS Country__deleted_at,Country.name AS Country__name,Country.capital_id AS Country__capital_id
	// FROM cities
	// LEFT JOIN countries Country ON
	//    Country.id = cities.country_id AND Country.deleted_at IS NULL
	// WHERE cities.name = "Paris" AND cities.deleted_at IS NULL

	if err != nil {
		log.Panicln(err)
	}

	log.Println("Cities named 'Paris' are:")
	for i, city := range cities {
		fmt.Printf("\t%v: %+v with country: ", i+1, city)

		cityCountry, err := city.GetCountry()
		if err != nil {
			log.Panicln(err)
		}

		fmt.Printf("%+v\n", cityCountry)
	}
}
