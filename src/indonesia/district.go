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

// District district of indonesia
type District struct {
	csvFile *os.File
}

// GetDistricts get all district of indonesia
func (ds *District) GetDistricts() ([]*src.District, error) {
	var districts []*src.District

	reader := csv.NewReader(bufio.NewReader(ds.csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		id, _ := strconv.Atoi(line[0])
		pid, _ := strconv.Atoi(line[1])
		districts = append(districts, &src.District{
			ID:       uint(id),
			Name:     line[2],
			ParentID: uint(pid),
		})
	}

	return districts, nil
}

// NewDistrict instance of indonesia district
func NewDistrict(filePath string) *District {
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return &District{csvFile}
}
