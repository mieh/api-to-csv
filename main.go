package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mieh/api-to-csv/output"
)

func main() {
	// Define command-line flags
	urlFlag := flag.String("url", "", "API URL")
	tokenFlag := flag.String("token", "", "Bearer Token")
	outputTypeFlag := flag.String("outputType", "toConsole", "Output Type ('toConsole', 'toFile', or 'toSheet')")
	outputPathFlag := flag.String("outputPath", "", "Output Path (Required for 'toFile' output type)")
	outputFileNameFlag := flag.String("outputFileName", "output.csv", "Output File Name (Used for 'toFile' output type)")
	spreadsheetIdFlag := flag.String("spreadsheetId", "", "Google Sheets Spreadsheet ID (Used for 'toSheet' output type)")
	sheetNameFlag := flag.String("sheetName", "Sheet1", "Sheet Name (Used for 'toSheet' output type)")

	flag.Parse()

	// Ensure that required flags are provided
	if *urlFlag == "" {
		log.Fatal("API URL is required")
	}
	if *tokenFlag == "" {
		log.Fatal("Bearer Token is required")
	}

	// Make API request and get response body
	responseBody, err := makeAPIRequest(*urlFlag, *tokenFlag)
	if err != nil {
		log.Fatalf("Error making API request: %v", err)
	}

	// Convert response body to CSV string
	csvString, err := convertToCSV(responseBody)
	if err != nil {
		log.Fatalf("Error converting response body to CSV: %v", err)
	}

	// Output the CSV string based on the outputType flag
	switch *outputTypeFlag {
	case "toConsole":
		err = output.ToConsole()(csvString)
		if err != nil {
			log.Fatalf("Error outputting CSV to console: %v", err)
		}
	case "toFile":
		if *outputPathFlag == "" {
			log.Fatal("Output Path is required for 'toFile' output type")
		}

		err = output.ToFile(csvString, *outputPathFlag, *outputFileNameFlag)
		if err != nil {
			log.Fatalf("Error outputting CSV to file: %v", err)
		}
	case "toSheet":
		if *spreadsheetIdFlag == "" {
			log.Fatal("Google Sheets Spreadsheet ID is required for 'toSheet' output type")
		}

		err = output.ToSheet(*spreadsheetIdFlag, *sheetNameFlag)(csvString)
		if err != nil {
			log.Fatalf("Error outputting CSV to Google Sheets: %v", err)
		}
	default:
		log.Fatalf("Invalid output type: %s", *outputTypeFlag)
	}

	fmt.Println("CSV conversion and output successful!")
}
