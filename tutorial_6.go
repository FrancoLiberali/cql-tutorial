//go:build tutorial_6
// +build tutorial_6

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-tutorial/conditions"
	"github.com/FrancoLiberali/cql-tutorial/models"
)

// Target: get all cities whose name is 'Paris' and that are the capital of their country
func tutorial(db *cql.DB) {
	cities, err := cql.Query[models.City](
		context.Background(),
		db,
		conditions.City.Name.Is().Eq(cql.String("Paris")),
		conditions.City.Country(
			conditions.Country.CapitalID.Is().Eq(conditions.City.ID),
		),
	).Find()

	// SQL executed:
	// SELECT cities.* FROM cities
	// INNER JOIN countries Country ON
	//    Country.id = cities.country_id AND Country.capital_id = cities.id
	// WHERE cities.name = "Paris"

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("--------------------------")
	fmt.Println("Cities named 'Paris' that are the capital of their country are:")
	for i, city := range cities {
		fmt.Printf("\t%v: %+v\n", i+1, city)
	}
}
