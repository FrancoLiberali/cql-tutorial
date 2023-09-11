//go:build tutorial_2
// +build tutorial_2

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

// Target: get all cities whose name is 'Paris' and its population is greater than 1000000
func tutorial(db *gorm.DB, shutdowner fx.Shutdowner) {
	cities, err := orm.Query[models.City](
		db,
		conditions.City.Name.Is().Eq("Paris"),
		conditions.City.Population.Is().Gt(1000000),
	).Find()

	// SQL executed:
	// SELECT cities.* FROM cities
	// WHERE cities.name = "Paris" AND cities.population > 1000000 AND cities.deleted_at IS NULL

	if err != nil {
		log.Panicln(err)
	}

	log.Println("Cities named 'Paris' with a population bigger than 1.000.000 are:")
	for i, city := range cities {
		fmt.Printf("\t%v: %+v\n", i+1, city)
	}

	shutdowner.Shutdown()
}
