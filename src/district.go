package src

// District entity
type District struct {
	ID       string
	Name     string
	ParentID string
}

// DistrictContract for contract to district
type DistrictContract interface {
	GetDistricts() ([]*District, error)
}
