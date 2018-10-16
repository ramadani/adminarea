package indonesia_test

import (
	"testing"

	"github.com/ramadani/adminarea/src/indonesia"
	"github.com/stretchr/testify/assert"
)

func TestGetDistricts(t *testing.T) {
	district := indonesia.NewDistrict("../../resources/indonesia/districts.csv")
	districts, err := district.GetDistricts()

	assert.Nil(t, err)
	assert.Equal(t, 7215, len(districts))
}
