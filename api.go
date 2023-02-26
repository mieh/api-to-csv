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

	var result []Response
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
