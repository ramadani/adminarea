package src

// Village entity
type Village struct {
	ID       string
	Name     string
	ParentID string
}

// VillageContract for contract to village
type VillageContract interface {
	GetVillages() ([]*Village, error)
}
