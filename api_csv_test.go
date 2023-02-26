package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

// Mock HTTP Client
type MockHTTPClient struct {
	Response *http.Response
	Err      error
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.Response, m.Err
}

func TestMakeAPIRequest(t *testing.T) {
	// Test with successful API response
	expectedResponse := "id,name,age\n1,John,30\n2,Jane,25\n"
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString(expectedResponse)),
	}
	client := &MockHTTPClient{Response: mockResponse, Err: nil}

	responseBody, err := makeAPIRequest("https://example.com/api/users", "Bearer ABC123", WithHTTPClient(client))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if responseBody != expectedResponse {
		t.Errorf("Unexpected response body: got %q, want %q", responseBody, expectedResponse)
	}

	// Test with non-200 response
	mockResponse = &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       ioutil.NopCloser(bytes.NewBufferString("")),
	}
	client = &MockHTTPClient{Response: mockResponse, Err: nil}

	responseBody, err = makeAPIRequest("https://example.com/api/users", "Bearer ABC123", WithHTTPClient(client))
	if err == nil {
		t.Errorf("Expected error, but got none")
	}
	if responseBody != "" {
		t.Errorf("Unexpected response body: got %q, want \"\"", responseBody)
	}

	// Test with HTTP client error
	client = &MockHTTPClient{Response: nil, Err: http.ErrAbortHandler}

	responseBody, err = makeAPIRequest("https://example.com/api/users", "Bearer ABC123", WithHTTPClient(client))
	if err == nil {
		t.Errorf("Expected error, but got none")
	}
	if responseBody != "" {
		t.Errorf("Unexpected response body: got %q, want \"\"", responseBody)
	}
}

func TestConvertToCSV(t *testing.T) {
	// Test with valid JSON
	jsonString := `[
		{"id": 1, "name": "John", "age": 30},
		{"id": 2, "name": "Jane", "age": 25}
	]`
	expectedCSV := "id,name,age\n1,John,30\n2,Jane,25\n"

	csvString, err := convertToCSV(jsonString)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if csvString != expectedCSV {
		t.Errorf("Unexpected CSV string: got %q, want %q", csvString, expectedCSV)
	}

	// Test with invalid JSON
	jsonString = `{id: 1, name: "John", age: 30}`
	_, err = convertToCSV(jsonString)
	if err == nil {
		t.Errorf("Expected error, but got none")
	}
}

func convertToCSV(jsonString string) {
	panic("unimplemented")
}
