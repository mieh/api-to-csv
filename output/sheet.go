package output

import (
	"context"
	"fmt"
	"log"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func ToSheet(spreadsheetId string, sheetName string) func(string) error {
	return func(csv string) error {
		ctx := context.Background()

		service, err := sheets.NewService(ctx, option.WithScopes(sheets.SpreadsheetsScope))
		if err != nil {
			log.Fatalf("Unable to retrieve Sheets client: %v", err)
		}

		rangeData := &sheets.ValueRange{
			Values: [][]interface{}{},
		}

		rows := strings.Split(csv, "\n")
		for _, row := range rows {
			if row != "" {
				cols := strings.Split(row, ",")
				var rowData []interface{}
				for _, col := range cols {
					rowData = append(rowData, col)
				}
				rangeData.Values = append(rangeData.Values, rowData)
			}
		}

		writeRange := fmt.Sprintf("%s!A1", sheetName)

		_, err = service.Spreadsheets.Values.Update(spreadsheetId, writeRange, rangeData).
			ValueInputOption("USER_ENTERED").
			Context(ctx).
			Do()
		if err != nil {
			log.Fatalf("Unable to write CSV data to sheet: %v", err)
		}

		return nil
	}
}
