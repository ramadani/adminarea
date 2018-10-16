package indonesia_test

import (
	"testing"

	"github.com/ramadani/adminarea/src/indonesia"
	"github.com/stretchr/testify/assert"
)

func TestGetRegencies(t *testing.T) {
	regency := indonesia.NewRegency("../../resources/indonesia/regencies.csv")
	regencies, err := regency.GetRegencies()

	assert.Nil(t, err)
	assert.Equal(t, 514, len(regencies))
}
