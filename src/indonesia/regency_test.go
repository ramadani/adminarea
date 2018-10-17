package indonesia_test

import (
	"log"
	"testing"

	idres "github.com/ramadani/adminarea/resources/indonesia"
	"github.com/ramadani/adminarea/src"

	"github.com/ramadani/adminarea/src/indonesia"
	"github.com/stretchr/testify/assert"
)

func TestGetRegencies(t *testing.T) {
	contents, err := idres.Asset("regencies.csv")
	if err != nil {
		log.Fatal(err)
	}

	countryRepo := src.NewCountryRepository()
	id, _ := countryRepo.GetCountry("ID")

	regency := indonesia.NewRegency(id.ID, contents)
	regencies, err := regency.GetRegencies()

	assert.Nil(t, err)
	assert.Equal(t, 514, len(regencies))
}
