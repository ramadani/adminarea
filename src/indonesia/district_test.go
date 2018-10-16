package indonesia_test

import (
	"log"
	"testing"

	idres "github.com/ramadani/adminarea/resources/indonesia"
	"github.com/ramadani/adminarea/src/indonesia"
	"github.com/stretchr/testify/assert"
)

func TestGetDistricts(t *testing.T) {
	contents, err := idres.Asset("districts.csv")
	if err != nil {
		log.Fatal(err)
	}

	district := indonesia.NewDistrict(contents)
	districts, err := district.GetDistricts()

	assert.Nil(t, err)
	assert.Equal(t, 7215, len(districts))
}
