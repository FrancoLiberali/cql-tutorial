package models

import (
	"github.com/FrancoLiberali/cql/model"
)

type Country struct {
	model.UUIDModel

	Name string

	Capital   *City // Country HasOne Capital (Country 1 -> 1 Capital)
	CapitalID *model.UUID

	Cities *[]City // Country HasMany City (Country 1 -> 1..* City)
}

type City struct {
	model.UUIDModel

	Name       string
	Population int

	Country   *Country
	CountryID *model.UUID // City HasOne Country (Country 1 <- 1..* City)
}
