// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package conditions

import (
	models "github.com/ditrit/badaas-orm-tutorial/models"
	orm "github.com/ditrit/badaas/orm"
	condition "github.com/ditrit/badaas/orm/condition"
	model "github.com/ditrit/badaas/orm/model"
	query "github.com/ditrit/badaas/orm/query"
	"reflect"
	"time"
)

var cityType = reflect.TypeOf(*new(models.City))

func (cityConditions cityConditions) IdIs() orm.FieldIs[models.City, model.UUID] {
	return orm.FieldIs[models.City, model.UUID]{FieldID: cityConditions.ID}
}
func (cityConditions cityConditions) CreatedAtIs() orm.FieldIs[models.City, time.Time] {
	return orm.FieldIs[models.City, time.Time]{FieldID: cityConditions.CreatedAt}
}
func (cityConditions cityConditions) UpdatedAtIs() orm.FieldIs[models.City, time.Time] {
	return orm.FieldIs[models.City, time.Time]{FieldID: cityConditions.UpdatedAt}
}
func (cityConditions cityConditions) DeletedAtIs() orm.FieldIs[models.City, time.Time] {
	return orm.FieldIs[models.City, time.Time]{FieldID: cityConditions.DeletedAt}
}
func (cityConditions cityConditions) NameIs() orm.StringFieldIs[models.City] {
	return orm.StringFieldIs[models.City]{FieldIs: orm.FieldIs[models.City, string]{FieldID: cityConditions.Name}}
}
func (cityConditions cityConditions) PopulationIs() orm.FieldIs[models.City, int] {
	return orm.FieldIs[models.City, int]{FieldID: cityConditions.Population}
}
func (cityConditions cityConditions) Country(conditions ...condition.Condition[models.Country]) condition.JoinCondition[models.City] {
	return condition.NewJoinCondition[models.City, models.Country](conditions, "Country", "CountryID", cityConditions.Preload(), "ID")
}
func (cityConditions cityConditions) PreloadCountry() condition.JoinCondition[models.City] {
	return cityConditions.Country(Country.Preload())
}
func (cityConditions cityConditions) CountryIdIs() orm.FieldIs[models.City, model.UUID] {
	return orm.FieldIs[models.City, model.UUID]{FieldID: cityConditions.CountryID}
}

type cityConditions struct {
	ID         query.FieldIdentifier[model.UUID]
	CreatedAt  query.FieldIdentifier[time.Time]
	UpdatedAt  query.FieldIdentifier[time.Time]
	DeletedAt  query.FieldIdentifier[time.Time]
	Name       query.FieldIdentifier[string]
	Population query.FieldIdentifier[int]
	CountryID  query.FieldIdentifier[model.UUID]
}

var City = cityConditions{
	CountryID: query.FieldIdentifier[model.UUID]{
		Field:     "CountryID",
		ModelType: cityType,
	},
	CreatedAt: query.FieldIdentifier[time.Time]{
		Field:     "CreatedAt",
		ModelType: cityType,
	},
	DeletedAt: query.FieldIdentifier[time.Time]{
		Field:     "DeletedAt",
		ModelType: cityType,
	},
	ID: query.FieldIdentifier[model.UUID]{
		Field:     "ID",
		ModelType: cityType,
	},
	Name: query.FieldIdentifier[string]{
		Field:     "Name",
		ModelType: cityType,
	},
	Population: query.FieldIdentifier[int]{
		Field:     "Population",
		ModelType: cityType,
	},
	UpdatedAt: query.FieldIdentifier[time.Time]{
		Field:     "UpdatedAt",
		ModelType: cityType,
	},
}

func (cityConditions cityConditions) Preload() condition.Condition[models.City] {
	return condition.NewPreloadCondition[models.City](cityConditions.ID, cityConditions.CreatedAt, cityConditions.UpdatedAt, cityConditions.DeletedAt, cityConditions.Name, cityConditions.Population, cityConditions.CountryID)
}
func (cityConditions cityConditions) PreloadRelations() []condition.Condition[models.City] {
	return []condition.Condition[models.City]{cityConditions.PreloadCountry()}
}
