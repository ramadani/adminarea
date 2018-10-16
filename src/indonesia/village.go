package indonesia

import (
	"bytes"
	"encoding/csv"
	"io"
	"strconv"

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

		id, _ := strconv.Atoi(line[0])
		pid, _ := strconv.Atoi(line[1])
		villages = append(villages, &src.Village{
			ID:       uint(id),
			Name:     line[2],
			ParentID: uint(pid),
		})
	}

	return villages, nil
}

// NewVillage instance of indonesia village
func NewVillage(contents []byte) *Village {
	return &Village{contents}
}
