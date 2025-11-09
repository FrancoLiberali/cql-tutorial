//go:build tutorial_4
// +build tutorial_4

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-tutorial/conditions"
	"github.com/FrancoLiberali/cql-tutorial/models"
)

// Target: get all cities whose name is 'Paris' and that the country to which they belong is called 'France'.
func tutorial(db *cql.DB) {
	parisFrance, err := cql.Query[models.City](
		context.Background(),
		db,
		conditions.City.Name.Is().Eq(cql.String("Paris")),
		conditions.City.Country(
			conditions.Country.Name.Is().Eq(cql.String("France")),
		),
	).FindOne()

	// SQL executed:
	// SELECT cities.* FROM cities
	// INNER JOIN countries Country ON
	//    Country.id = cities.country_id AND Country.name = "France"
	// WHERE cities.name = "Paris"

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("--------------------------")
	fmt.Printf("City named 'Paris' in 'France' is: %+v\n", parisFrance)
}
