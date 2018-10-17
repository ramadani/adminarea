package indonesia

import (
	"bytes"
	"encoding/csv"
	"io"

	"github.com/ramadani/adminarea/src"
)

// Province province of indonesia
type Province struct {
	countryID string
	contents  []byte
}

// All get all province
func (pr *Province) All() ([]*src.Province, error) {
	var provinces []*src.Province

	reader := csv.NewReader(bytes.NewReader(pr.contents))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		provinces = append(provinces, &src.Province{
			ID:       src.AreaID(pr.countryID, line[0]),
			Name:     line[1],
			ParentID: pr.countryID,
		})
	}

	return provinces, nil
}

// NewProvince instance of indonesia provinces
func NewProvince(countryID string, contents []byte) *Province {
	return &Province{countryID, contents}
}
