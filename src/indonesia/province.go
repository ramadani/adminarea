package indonesia

import (
	"bytes"
	"encoding/csv"
	"io"
	"strconv"

	"github.com/ramadani/adminarea/src"
)

// Province province of indonesia
type Province struct {
	countryID uint
	contents  []byte
}

// GetProvinces get all province
func (pr *Province) GetProvinces() ([]*src.Province, error) {
	var provinces []*src.Province

	reader := csv.NewReader(bytes.NewReader(pr.contents))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		id, _ := strconv.Atoi(line[0])
		provinces = append(provinces, &src.Province{
			ID:       uint(id),
			Name:     line[1],
			ParentID: pr.countryID,
		})
	}

	return provinces, nil
}

// NewProvince instance of indonesia provinces
func NewProvince(countryID uint, contents []byte) *Province {
	return &Province{countryID, contents}
}
