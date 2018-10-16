package src

// Regency entity
type Regency struct {
	ID       uint
	Name     string
	ParentID uint
}

// RegencyContract for contract to province
type RegencyContract interface {
	GetRegencies() []*Regency
}
