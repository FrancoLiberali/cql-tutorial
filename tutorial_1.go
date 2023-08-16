//go:build tutorial_1
// +build tutorial_1

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

// Target: get all cities whose name is 'Paris'
// SQL executed: SELECT cities.* FROM cities WHERE cities.name = "Paris" AND cities.deleted_at IS NULL
func tutorial(db *gorm.DB, shutdowner fx.Shutdowner) {
	cities, err := orm.NewQuery[models.City](
		db,
		conditions.City.NameIs().Eq("Paris"),
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

	shutdowner.Shutdown()
}
