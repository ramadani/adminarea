package indonesia

import (
	"bytes"
	"encoding/csv"
	"io"
	"strconv"

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

		id, _ := strconv.Atoi(line[0])
		pid, _ := strconv.Atoi(line[1])
		regencies = append(regencies, &src.Regency{
			ID:       uint(id),
			Name:     line[2],
			ParentID: uint(pid),
		})
	}

	return regencies, nil
}

// NewRegency instance of indonesia regency
func NewRegency(contents []byte) *Regency {
	return &Regency{contents}
}
