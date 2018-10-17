package indonesia

import (
	"bytes"
	"encoding/csv"
	"io"

	"github.com/ramadani/adminarea/src"
)

// District district of indonesia
type District struct {
	countryID string
	contents  []byte
}

// All get all district of indonesia
func (ds *District) All() ([]*src.District, error) {
	var districts []*src.District

	reader := csv.NewReader(bytes.NewReader(ds.contents))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		districts = append(districts, &src.District{
			ID:       src.AreaID(ds.countryID, line[0]),
			Name:     line[2],
			ParentID: src.AreaID(ds.countryID, line[1]),
		})
	}

	return districts, nil
}

// NewDistrict instance of indonesia district
func NewDistrict(countryID string, contents []byte) *District {
	return &District{countryID, contents}
}
