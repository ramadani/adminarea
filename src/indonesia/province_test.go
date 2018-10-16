package indonesia_test

import (
	"log"
	"testing"

	"github.com/ramadani/adminarea/src/indonesia"
	"github.com/stretchr/testify/assert"
)

func TestGetProvinces(t *testing.T) {
	province := indonesia.NewProvince("../../resources/indonesia/provinces.csv")
	provinces, err := province.GetProvinces()
	log.Println(provinces)
	assert.Nil(t, err)
	assert.Equal(t, 34, len(provinces))
}
