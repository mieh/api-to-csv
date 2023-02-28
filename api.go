package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Field1 string `json:"id"`
	Field2 string `json:"payment_no"`
	Field3 string `json:"transaction_date"`
	Field4 string `json:"total_collected"`
}

func GetApiData(apiUrl string, bearerToken string) ([]Response, error) {
	var result []Response

	for apiUrl != "" {
		// Make API request
		client := &http.Client{}
		req, err := http.NewRequest("GET", apiUrl, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+bearerToken)

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// Decode response body
		var response struct {
			Results []Response `json:"results"`
			NextUrl string     `json:"nextUrl"`
		}

		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return nil, err
		}

		// Append results to the overall result
		result = append(result, response.Results...)

		// Update apiUrl for the next page
		apiUrl = response.NextUrl
	}

	return result, nil
}
