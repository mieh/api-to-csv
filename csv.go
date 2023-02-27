package main

import (
	"bytes"
	"encoding/csv"
)

func ConvertToCsv(data []Response) (string, error) {
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)

	headers := []string{"field1", "field2", "field3", "field4"}
	err := w.Write(headers)
	if err != nil {
		return "", err
	}

	for _, d := range data {
		row := []string{d.Field1, d.Field2}
		err := w.Write(row)
		if err != nil {
			return "", err
		}
	}

	w.Flush()

	return buf.String(), nil
}
