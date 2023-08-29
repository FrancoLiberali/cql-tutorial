package models

import (
	"github.com/ditrit/badaas/orm/model"
)

type Country struct {
	model.UUIDModel

	Name string

	Capital   *City // Country HasOne Capital (Country 1 -> 1 Capital)
	CapitalID *model.UUID
}

type City struct {
	model.UUIDModel

	Name       string
	Population int

	Country   *Country
	CountryID *model.UUID // City HasOne Country (Country 1 <- 1..* City)
}
