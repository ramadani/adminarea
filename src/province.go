package src

// Province entity
type Province struct {
	ID   uint
	Name string
}

// ProvinceContract for contract to province
type ProvinceContract interface {
	GetProvinces() []*Province
}
