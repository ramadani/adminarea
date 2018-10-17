package adminarea_test

import (
	"testing"

	"github.com/ramadani/adminarea"
	"github.com/stretchr/testify/assert"
)

func TestIndonesiaGetProvinces(t *testing.T) {
	assert := assert.New(t)
	idAdminArea := adminarea.New("ID")
	provinces, err := idAdminArea.GetProvinces()

	assert.Nil(err)
	assert.Equal(34, len(provinces))
}

func TestIndonesiaGetRegencies(t *testing.T) {
	assert := assert.New(t)
	idAdminArea := adminarea.New("ID")
	regencies, err := idAdminArea.GetRegencies()

	assert.Nil(err)
	assert.Equal(514, len(regencies))
}

func TestIndonesiaGetDistricts(t *testing.T) {
	assert := assert.New(t)
	idAdminArea := adminarea.New("ID")
	districts, err := idAdminArea.GetDistricts()

	assert.Nil(err)
	assert.Equal(7215, len(districts))
}

func TestIndonesiaGetVillages(t *testing.T) {
	assert := assert.New(t)
	idAdminArea := adminarea.New("ID")
	villages, err := idAdminArea.GetVillages()

	assert.Nil(err)
	assert.Equal(80534, len(villages))
}
