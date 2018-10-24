package src

import (
	"errors"
	"fmt"
)

// Country entity
type Country struct {
	ID   string
	Name string
}

// CountryRepository country repository
type CountryRepository struct {
	countries []*Country
}

// FindByID by id
func (c *CountryRepository) FindByID(id string) (*Country, error) {
	for _, country := range c.countries {
		if id == country.ID {
			return country, nil
		}
	}

	return nil, errors.New("Country is not available")
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
