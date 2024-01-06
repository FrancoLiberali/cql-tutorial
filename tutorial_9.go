//go:build tutorial_9
// +build tutorial_9

package main

import (
	"fmt"
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-tutorial/conditions"
	"github.com/FrancoLiberali/cql-tutorial/models"
	"gorm.io/gorm"
)

// Target: obtain all the countries that have a city called 'Paris'
func tutorial(db *gorm.DB) {
	countries, err := cql.Query[models.Country](
		db,
		conditions.Country.Cities.Any(
			conditions.City.Name.Is().Eq("Paris"),
		),
	).Find()

	// SQL executed:
	// SELECT countries.* FROM countries
	// WHERE (EXISTS (
	//     SELECT(1) FROM cities
	//     WHERE cities.country_id = countries.id AND
	//           cities.name = "Paris" AND
	//           cities.deleted_at IS NULL
	// )) AND countries.deleted_at IS NULL

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("--------------------------")
	fmt.Println("Countries that have a city called 'Paris' are:")
	for i, country := range countries {
		fmt.Printf("\t%v: %+v\n", i+1, country)
	}
}
