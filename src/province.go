package src

// Province entity
type Province struct {
	ID       string
	Name     string
	ParentID string
}

// ProvinceContract for contract to province
type ProvinceContract interface {
	All() ([]*Province, error)
}
