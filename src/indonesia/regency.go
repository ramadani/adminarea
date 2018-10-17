package indonesia

import (
	"bytes"
	"encoding/csv"
	"io"

	"github.com/ramadani/adminarea/src"
)

// Regency regency of indonesia
type Regency struct {
	contents []byte
}

// GetRegencies get all regency of indonesia
func (rg *Regency) GetRegencies() ([]*src.Regency, error) {
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
			ID:       line[0],
			Name:     line[2],
			ParentID: line[1],
		})
	}

	return regencies, nil
}

// NewRegency instance of indonesia regency
func NewRegency(contents []byte) *Regency {
	return &Regency{contents}
}
