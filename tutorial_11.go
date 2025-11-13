//go:build tutorial_11
// +build tutorial_11

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-tutorial/conditions"
	"github.com/FrancoLiberali/cql-tutorial/models"
)

type CityAndCountryNames struct {
	CityName    string
	CountryName string
}

// Target: obtain only the city and country name of Paris, France
func tutorial(db *cql.DB) {
	results, err := cql.Select(
		cql.Query[models.City](
			context.Background(),
			db,
			conditions.City.Name.Is().Eq(cql.String("Paris")),
			conditions.City.Country(
				conditions.Country.Name.Is().Eq(cql.String("France")),
			),
		),
		cql.ValueInto(conditions.City.Name, func(value string, result *CityAndCountryNames) {
			result.CityName = value
		}),
		cql.ValueInto(conditions.Country.Name, func(value string, result *CityAndCountryNames) {
			result.CountryName = value
		}),
	)

	// SQL executed:
	// SELECT cities.name, Country.name
	// FROM cities
	// INNER JOIN countries Country ON
	//    Country.id = cities.country_id AND Country.name = "France"
	// WHERE cities.name = "Paris"

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("--------------------------")
	fmt.Printf("City named 'Paris' in 'France' is: %+v\n", results)
}
