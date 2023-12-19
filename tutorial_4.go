//go:build tutorial_4
// +build tutorial_4

package main

import (
	"fmt"
	"log"

	"github.com/ditrit/badaas-orm-tutorial/conditions"
	"github.com/ditrit/badaas-orm-tutorial/models"
	"github.com/ditrit/badaas/orm"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Target: get all cities whose name is 'Paris' and that the country to which they belong is called 'France'.
func tutorial(db *gorm.DB, shutdowner fx.Shutdowner) {
	parisFrance, err := orm.Query[models.City](
		db,
		conditions.City.Name.Is().Eq("Paris"),
		conditions.City.Country(
			conditions.Country.Name.Is().Eq("France"),
		),
	).FindOne()

	// SQL executed:
	// SELECT cities.* FROM cities
	// INNER JOIN countries Country ON
	//    Country.id = cities.country_id AND Country.name = "France" AND Country.deleted_at IS NULL
	// WHERE cities.name = "Paris" AND cities.deleted_at IS NULL

	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("City named 'Paris' in 'France' is: %+v\n", parisFrance)

	shutdowner.Shutdown()
}
