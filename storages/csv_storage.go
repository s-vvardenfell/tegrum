package storages

import (
	"encoding/csv"
	"fmt"
	"io"
)

type CsvStorage struct {
}

func (s *CsvStorage) Store(w io.Writer, data []string) error {
	writer := csv.NewWriter(w)
	defer writer.Flush()
	if err := writer.Write(data); err != nil {
		return fmt.Errorf("error while writing data, %v", err)
	}
	return nil
}

func (s *CsvStorage) Retrieve(r io.Reader, index string) ([]string, error) {
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = 3
	reader.Comment = '#'

	for {
		record, err := reader.Read()
		if err != nil {
			return nil, fmt.Errorf("error while fetching records, %v", err)
		}

		for _, r := range record {
			if r == index {
				return record, nil
			}
		}
	}
}
