//go:build tutorial_10
// +build tutorial_10

package main

import (
	"context"
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-tutorial/conditions"
	"github.com/FrancoLiberali/cql-tutorial/models"
)

// Target: verify that cql is compile-time safe
// Will not compile
func tutorial(db *cql.DB) {
	_, err := cql.Query[models.City](
		context.Background(),
		db,
		conditions.Country.Name.Is().Eq(cql.String("Paris")),
	).Find()

	if err != nil {
		log.Panicln(err)
	}
}
