package src

// Country entity
type Country struct {
	ID   uint
	Name string
}

// CountryContract for contract to country
type CountryContract interface {
	GetCountry(id uint) *Country
}

// CountryRepository country repository
type CountryRepository struct {
	countries []*Country
}

// GetCountry by id
func (c *CountryRepository) GetCountry(id uint) *Country {
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
			&Country{ID: 360, Name: "Indonesia"},
		},
	}
}
