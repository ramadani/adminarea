package src

// Regency entity
type Regency struct {
	ID       string
	Name     string
	ParentID string
}

// RegencyContract for contract to regency
type RegencyContract interface {
	All() ([]*Regency, error)
}
