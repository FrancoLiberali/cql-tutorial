//go:build tutorial_9
// +build tutorial_9

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-tutorial/conditions"
	"github.com/FrancoLiberali/cql-tutorial/models"
)

// Target: create Rennes in France and then delete it
func tutorial(db *cql.DB) {
	// get country called France
	france, err := cql.Query[models.Country](
		context.Background(),
		db,
		conditions.Country.Name.Is().Eq(cql.String("France")),
	).FindOne()

	if err != nil {
		log.Panicln(err)
	}

	// create Rennes
	rennes := models.City{
		CountryID:  france.ID,
		Name:       "Rennes",
		Population: 215366,
	}

	inserted, err := cql.Insert(context.Background(), db, &rennes).Exec()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Inserted %v city\n", inserted)

	// delete city called Rennes
	deleted, err := cql.Delete[models.City](
		context.Background(),
		db,
		conditions.City.Name.Is().Eq(cql.String("Rennes")),
	).Exec()

	// SQL executed:
	// DELETE FROM cities
	// WHERE cities.name = "Rennes"

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("--------------------------")
	fmt.Printf("Deleted %v city\n", deleted)
}
