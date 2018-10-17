package src

// Province entity
type Province struct {
	ID       uint
	Name     string
	ParentID uint
}

// ProvinceContract for contract to province
type ProvinceContract interface {
	GetProvinces() ([]*Province, error)
}
