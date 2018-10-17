package indonesia

import (
	"bytes"
	"encoding/csv"
	"io"

	"github.com/ramadani/adminarea/src"
)

// Village village of indonesia
type Village struct {
	countryID string
	contents  []byte
}

// All get all village of indonesia
func (vl *Village) All() ([]*src.Village, error) {
	var villages []*src.Village

	reader := csv.NewReader(bytes.NewReader(vl.contents))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		villages = append(villages, &src.Village{
			ID:       src.AreaID(vl.countryID, line[0]),
			Name:     line[2],
			ParentID: src.AreaID(vl.countryID, line[1]),
		})
	}

	return villages, nil
}

// NewVillage instance of indonesia village
func NewVillage(countryID string, contents []byte) *Village {
	return &Village{countryID, contents}
}
