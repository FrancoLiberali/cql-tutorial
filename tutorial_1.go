//go:build tutorial_1
// +build tutorial_1

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-tutorial/conditions"
	"github.com/FrancoLiberali/cql-tutorial/models"
)

// Target: get all cities whose name is 'Paris'
// SQL executed: SELECT cities.* FROM cities WHERE cities.name = "Paris"
func tutorial(db *cql.DB) {
	cities, err := cql.Query[models.City](
		context.Background(),
		db,
		conditions.City.Name.Is().Eq(cql.String("Paris")),
	).Find()

	// SQL executed:
	// SELECT cities.* FROM cities
	// WHERE cities.name = "Paris"

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("--------------------------")
	fmt.Println("Cities named 'Paris' are:")
	for i, city := range cities {
		fmt.Printf("\t%v: %+v\n", i+1, city)
	}
}
