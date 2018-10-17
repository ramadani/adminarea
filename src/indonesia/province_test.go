package indonesia_test

import (
	"log"
	"testing"

	idres "github.com/ramadani/adminarea/resources/indonesia"
	"github.com/ramadani/adminarea/src"
	"github.com/ramadani/adminarea/src/indonesia"
	"github.com/stretchr/testify/assert"
)

func TestGetProvinces(t *testing.T) {
	contents, err := idres.Asset("provinces.csv")
	if err != nil {
		log.Fatal(err)
	}

	countryRepo := src.NewCountryRepository()
	id := countryRepo.GetCountry(360)

	province := indonesia.NewProvince(id.ID, contents)
	provinces, err := province.GetProvinces()

	assert.Nil(t, err)
	assert.Equal(t, 34, len(provinces))
}
