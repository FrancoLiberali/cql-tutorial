// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package models

import preload "github.com/ditrit/badaas/orm/preload"

func (m City) GetCountry() (*Country, error) {
	return preload.VerifyPointerLoaded[Country](m.CountryID, m.Country)
}
func (m Country) GetCapital() (*City, error) {
	return preload.VerifyPointerLoaded[City](m.CapitalID, m.Capital)
}
