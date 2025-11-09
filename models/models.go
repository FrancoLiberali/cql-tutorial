package models

import (
	"fmt"

	"github.com/FrancoLiberali/cql/model"
)

type Country struct {
	model.UIntModel

	Name string

	Capital   *City // Country HasOne Capital (Country 1 -> 1 Capital)
	CapitalID model.UIntID

	Cities *[]City // Country HasMany City (Country 1 -> 1..* City)
}

func (country Country) String() string {
	return fmt.Sprintf(
		"Country{ID: %v, Name: %s, CapitalID:%v, Capital:%s }",
		country.ID,
		country.Name,
		country.CapitalID,
		country.Capital,
	)
}

type City struct {
	model.UIntModel

	Name       string
	Population int

	Country   *Country
	CountryID model.UIntID `gorm:"not null"` // City HasOne Country (Country 1 <- 1..* City)
}

func (city City) String() string {
	return fmt.Sprintf(
		"City{ID: %v, Name: %s, Population: %v, CountryID:%v, Country:%v }",
		city.ID,
		city.Name,
		city.Population,
		city.CountryID,
		city.Country,
	)
}
