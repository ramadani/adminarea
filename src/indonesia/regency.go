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

// Regency regency of indonesia
type Regency struct {
	csvFile *os.File
}

// GetRegencies get all regency of indonesia
func (rg *Regency) GetRegencies() ([]*src.Regency, error) {
	var regencies []*src.Regency

	reader := csv.NewReader(bufio.NewReader(rg.csvFile))
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
func NewRegency(filePath string) *Regency {
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return &Regency{csvFile}
}
