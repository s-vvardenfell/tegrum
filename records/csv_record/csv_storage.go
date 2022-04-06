package csv_record

import (
	"encoding/csv"
	"fmt"
	"io"
)

type CsvRecorderRetriever struct {
}

func (s *CsvRecorderRetriever) Record(w io.Writer, data []string) error {
	writer := csv.NewWriter(w)
	defer writer.Flush()
	if err := writer.Write(data); err != nil {
		return fmt.Errorf("error while writing data, %v", err)
	}
	return nil
}

func (s *CsvRecorderRetriever) Retrieve(r io.Reader, index string) ([]string, error) {
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = 3
	reader.Comment = '#'

	records := make([][]string, 0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error while fetching records, %v", err)
		}
		records = append(records, record)
	}

	for i := len(records) - 1; i >= 0; i-- {
		for _, r := range records[i] {
			if r == index {
				return records[i], nil
			}
		}
	}
	return nil, fmt.Errorf("record with index %s not found\n", index)
}
