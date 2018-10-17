package src

import (
	"fmt"
)

// Country entity
type Country struct {
	ID   string
	Name string
}

// CountryContract for contract to country
type CountryContract interface {
	GetCountry(id string) *Country
}

// CountryRepository country repository
type CountryRepository struct {
	countries []*Country
}

// GetCountry by id
func (c *CountryRepository) GetCountry(id string) *Country {
	for _, country := range c.countries {
		if id == country.ID {
			return country
		}
	}

	return nil
}

// NewCountryRepository country repository
func NewCountryRepository() *CountryRepository {
	return &CountryRepository{
		[]*Country{
			&Country{ID: "ID", Name: "Indonesia"},
		},
	}
}

// AreaID to generate area id
func AreaID(id, subid string) string {
	return fmt.Sprintf("%s-%s", id, subid)
}
