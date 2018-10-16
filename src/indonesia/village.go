package indonesia

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/ramadani/adminarea/src"
)

// Village village of indonesia
type Village struct {
	csvFile *os.File
}

// GetVillages get all village of indonesia
func (vl *Village) GetVillages() ([]*src.Village, error) {
	var villages []*src.Village

	reader := csv.NewReader(bufio.NewReader(vl.csvFile))
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
func NewVillage(filePath string) *Village {
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return &Village{csvFile}
}
