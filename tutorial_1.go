//go:build tutorial_1
// +build tutorial_1

package main

import (
	"fmt"
	"log"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/cql-tutorial/conditions"
	"github.com/FrancoLiberali/cql-tutorial/models"
	"gorm.io/gorm"
)

// Target: get all cities whose name is 'Paris'
// SQL executed: SELECT cities.* FROM cities WHERE cities.name = "Paris" AND cities.deleted_at IS NULL
func tutorial(db *gorm.DB) {
	cities, err := cql.Query[models.City](
		db,
		conditions.City.Name.Is().Eq("Paris"),
	).Find()

	// SQL executed:
	// SELECT cities.* FROM cities
	// WHERE cities.name = "Paris" AND cities.deleted_at IS NULL

	if err != nil {
		log.Panicln(err)
	}

	log.Println("Cities named 'Paris' are:")
	for i, city := range cities {
		fmt.Printf("\t%v: %+v\n", i+1, city)
	}
}
