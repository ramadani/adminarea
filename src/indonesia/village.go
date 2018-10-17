package indonesia

import (
	"bytes"
	"encoding/csv"
	"io"

	"github.com/ramadani/adminarea/src"
)

// Village village of indonesia
type Village struct {
	contents []byte
}

// GetVillages get all village of indonesia
func (vl *Village) GetVillages() ([]*src.Village, error) {
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
			ID:       line[0],
			Name:     line[2],
			ParentID: line[1],
		})
	}

	return villages, nil
}

// NewVillage instance of indonesia village
func NewVillage(contents []byte) *Village {
	return &Village{contents}
}
