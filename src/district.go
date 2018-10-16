package src

// District entity
type District struct {
	ID       uint
	Name     string
	ParentID uint
}

// DistrictContract for contract to district
type DistrictContract interface {
	GetDistricts() ([]*District, error)
}
