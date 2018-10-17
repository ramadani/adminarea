package indonesia_test

import (
	"log"
	"testing"

	idres "github.com/ramadani/adminarea/resources/indonesia"
	"github.com/ramadani/adminarea/src"
	"github.com/ramadani/adminarea/src/indonesia"
	"github.com/stretchr/testify/assert"
)

func TestGetVillages(t *testing.T) {
	contents, err := idres.Asset("villages.csv")
	if err != nil {
		log.Fatal(err)
	}

	countryRepo := src.NewCountryRepository()
	id := countryRepo.GetCountry("ID")

	village := indonesia.NewVillage(id.ID, contents)
	villages, err := village.GetVillages()

	assert.Nil(t, err)
	assert.Equal(t, 80534, len(villages))
}
