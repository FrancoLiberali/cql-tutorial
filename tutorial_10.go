//go:build tutorial_10
// +build tutorial_10

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-tutorial/conditions"
	"github.com/FrancoLiberali/cql-tutorial/models"
)

// Target: obtain all the countries that have a city called 'Paris'
func tutorial(db *cql.DB) {
	countries, err := cql.Query[models.Country](
		context.Background(),
		db,
		conditions.Country.Cities.Any(
			conditions.City.Name.Is().Eq(cql.String("Paris")),
		),
	).Find()

	// SQL executed:
	// SELECT countries.* FROM countries
	// WHERE (EXISTS (
	//     SELECT(1) FROM cities
	//     WHERE cities.country_id = countries.id AND
	//           cities.name = "Paris"
	// ))

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("--------------------------")
	fmt.Println("Countries that have a city called 'Paris' are:")
	for i, country := range countries {
		fmt.Printf("\t%v: %+v\n", i+1, country)
	}
}
