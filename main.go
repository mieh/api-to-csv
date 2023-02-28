package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mieh/api-to-csv/output"
	"gopkg.in/yaml.v2"
)

type Config struct {
	URL            string `yaml:"url"`
	Token          string `yaml:"token"`
	OutputType     string `yaml:"outputType"`
	OutputPath     string `yaml:"outputPath"`
	OutputFileName string `yaml:"outputFileName"`
	SpreadsheetID  string `yaml:"spreadsheetId"`
	SheetName      string `yaml:"sheetName"`
}

func main() {
	// Define command-line flags
	configFlag := flag.String("config", "", "yml file name, consist of parameter value needed")

	flag.Parse()

	// Ensure that required flags are provided
	if *configFlag == "" {
		log.Fatal("config file name is required")
	}

	// Read in the YAML configuration file
	configFile, err := ioutil.ReadFile(*configFlag)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Parse the YAML configuration file into a Config struct
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	// Ensure that required parameters are provided
	if config.URL == "" {
		log.Fatal("API URL is required")
	}
	if config.Token == "" {
		log.Fatal("Bearer Token is required")
	}

	// Make API request and get response body
	responseBody, err := GetApiData(config.URL, config.Token)
	if err != nil {
		log.Fatalf("Error making API request: %v", err)
	}

	// Convert response body to CSV string
	csvString, err := ConvertToCsv(responseBody)
	if err != nil {
		log.Fatalf("Error converting response body to CSV: %v", err)
	}

	// Output the CSV string based on the outputType flag
	switch config.OutputType {
	case "toConsole":
		err = output.ToConsole(csvString)
		if err != nil {
			log.Fatalf("Error outputting CSV to console: %v", err)
		}
	case "toFile":
		if config.OutputPath == "" {
			log.Fatal("Output Path is required for 'toFile' output type")
		}

		err = output.ToFile(csvString, config.OutputPath, config.OutputFileName)
		if err != nil {
			log.Fatalf("Error outputting CSV to file: %v", err)
		}
	case "toSheet":
		if config.SpreadsheetID == "" {
			log.Fatal("Google Sheets Spreadsheet ID is required for 'toSheet' output type")
		}

		err = output.ToSheet(config.SpreadsheetID, config.SheetName)(csvString)
		if err != nil {
			log.Fatalf("Error outputting CSV to Google Sheets: %v", err)
		}
	default:
		log.Fatalf("Invalid output type: %s", config.OutputType)
	}

	fmt.Println("CSV conversion and output successful!")
}
