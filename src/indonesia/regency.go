package indonesia

import (
	"bytes"
	"encoding/csv"
	"io"

	"github.com/ramadani/adminarea/src"
)

// Regency regency of indonesia
type Regency struct {
	countryID string
	contents  []byte
}

// All get all regency of indonesia
func (rg *Regency) All() ([]*src.Regency, error) {
	var regencies []*src.Regency

	reader := csv.NewReader(bytes.NewReader(rg.contents))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		regencies = append(regencies, &src.Regency{
			ID:       src.AreaID(rg.countryID, line[0]),
			Name:     line[2],
			ParentID: src.AreaID(rg.countryID, line[1]),
		})
	}

	return regencies, nil
}

// NewRegency instance of indonesia regency
func NewRegency(countryID string, contents []byte) *Regency {
	return &Regency{countryID, contents}
}
