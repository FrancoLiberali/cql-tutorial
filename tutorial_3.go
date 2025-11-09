//go:build tutorial_3
// +build tutorial_3

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-tutorial/conditions"
	"github.com/FrancoLiberali/cql-tutorial/models"
)

// Target: get the city named 'Paris' with the largest population
func tutorial(db *cql.DB) {
	parisFrance, err := cql.Query[models.City](
		context.Background(),
		db,
		conditions.City.Name.Is().Eq(cql.String("Paris")),
	).Descending(
		conditions.City.Population,
	).Limit(1).FindOne()

	// SQL executed:
	// SELECT cities.* FROM cities
	// WHERE cities.name = "Paris"
	// ORDER BY cities.population DESC
	// LIMIT 1

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("--------------------------")
	fmt.Printf("City named 'Paris' with the largest population is: %+v\n", parisFrance)
}
