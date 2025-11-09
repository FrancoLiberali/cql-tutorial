//go:build tutorial_5
// +build tutorial_5

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-tutorial/conditions"
	"github.com/FrancoLiberali/cql-tutorial/models"
)

// Target: get all cities whose name is 'Paris' and preload its country
func tutorial(db *cql.DB) {
	cities, err := cql.Query[models.City](
		context.Background(),
		db,
		conditions.City.Name.Is().Eq(cql.String("Paris")),
		conditions.City.Country().Preload(),
	).Find()

	// SQL executed:
	// SELECT cities.*,
	//    Country.id AS Country__id,Country.name AS Country__name,Country.capital_id AS Country__capital_id
	// FROM cities
	// LEFT JOIN countries Country ON
	//    Country.id = cities.country_id
	// WHERE cities.name = "Paris"

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("--------------------------")
	fmt.Println("Cities named 'Paris' are:")
	for i, city := range cities {
		fmt.Printf("\t%v: %+v with country: ", i+1, city)

		cityCountry, err := city.GetCountry()
		if err != nil {
			log.Panicln(err)
		}

		fmt.Printf("%+v\n", cityCountry)
	}
}
