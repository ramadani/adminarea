package indonesia_test

import (
	"testing"

	"github.com/ramadani/adminarea/src/indonesia"
	"github.com/stretchr/testify/assert"
)

func TestGetVillages(t *testing.T) {
	village := indonesia.NewVillage("../../resources/indonesia/villages.csv")
	villages, err := village.GetVillages()

	assert.Nil(t, err)
	assert.Equal(t, 80534, len(villages))
}
