//go:build tutorial_6
// +build tutorial_6

package main

import (
	"fmt"
	"log"

	"github.com/ditrit/badaas-orm-tutorial/conditions"
	"github.com/ditrit/badaas-orm-tutorial/models"
	"github.com/ditrit/badaas/orm"
	"gorm.io/gorm"
)

// Target: get all cities whose name is 'Paris' and that are the capital of their country
func tutorial(db *gorm.DB) {
	cities, err := orm.Query[models.City](
		db,
		conditions.City.Name.Is().Eq("Paris"),
		conditions.City.Country(
			conditions.Country.CapitalID.Is().Dynamic().Eq(conditions.City.ID),
		),
	).Find()

	// SQL executed:
	// SELECT cities.* FROM cities
	// INNER JOIN countries Country ON
	//    Country.id = cities.country_id AND Country.capital_id = cities.id AND Country.deleted_at IS NULL
	// WHERE cities.name = "Paris" AND cities.deleted_at IS NULL

	if err != nil {
		log.Panicln(err)
	}

	log.Println("Cities named 'Paris' that are the capital of their country are:")
	for i, city := range cities {
		fmt.Printf("\t%v: %+v\n", i+1, city)
	}
}
