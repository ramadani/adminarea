package src

// Village entity
type Village struct {
	ID       uint
	Name     string
	ParentID uint
}

// VillageContract for contract to village
type VillageContract interface {
	GetVillages() ([]*Village, error)
}
