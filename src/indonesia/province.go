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

// Province province of indonesia
type Province struct {
	csvFile *os.File
}

// GetProvinces get all province
func (pr *Province) GetProvinces() ([]*src.Province, error) {
	var provinces []*src.Province

	reader := csv.NewReader(bufio.NewReader(pr.csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		id, _ := strconv.Atoi(line[0])
		provinces = append(provinces, &src.Province{
			ID:   uint(id),
			Name: line[1],
		})
	}

	return provinces, nil
}

// NewProvince instance of indonesia provinces
func NewProvince(filePath string) *Province {
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return &Province{csvFile}
}
