//go:build tutorial_9
// +build tutorial_9

package main

import (
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-tutorial/conditions"
	"github.com/FrancoLiberali/cql-tutorial/models"
	"gorm.io/gorm"
)

// Target: verify that cql is compile-time safe
// Will not compile
func tutorial(db *gorm.DB) {
	_, err := cql.Query[models.City](
		db,
		conditions.Country.Name.Is().Eq("Paris"),
	).Find()

	if err != nil {
		log.Panicln(err)
	}
}
