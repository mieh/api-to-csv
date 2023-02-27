package main

import (
	"encoding/json"
	"github.com/mieh/api-to-csv/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetApiData(t *testing.T) {
	// Set up a mock API server
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]Response{
			{
				Field1: "123",
				Field2: "abc",
				Field3: "2022-02-26",
				Field4: "100",
			},
			{
				Field1: "456",
				Field2: "def",
				Field3: "2022-02-27",
				Field4: "200",
			},
		})
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	// Make a request to the mock API server with a fake token
	token := "fake_token"
	response, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Error making API request: %v", err)
	}
	defer response.Body.Close()

	var result []Response
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		t.Fatalf("Error decoding API response: %v", err)
	}

	// Verify that the response matches the expected data
	expected := []Response{
		{
			Field1: "123",
			Field2: "abc",
			Field3: "2022-02-26",
			Field4: "100",
		},
		{
			Field1: "456",
			Field2: "def",
			Field3: "2022-02-27",
			Field4: "200",
		},
	}
	for i, r := range result {
		if r != expected[i] {
			t.Errorf("Unexpected response data at index %d: got %+v, expected %+v", i, r, expected[i])
		}
	}
}
