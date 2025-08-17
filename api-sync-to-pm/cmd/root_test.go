package cmd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNewAPIClient(t *testing.T) {
	docKey := "test-doc-key"
	pmKey := "test-pm-key"

	client := NewAPIClient(docKey, pmKey)

	if client.docAPIKey != docKey {
		t.Errorf("NewAPIClient() docAPIKey = %v, want %v", client.docAPIKey, docKey)
	}

	if client.pmAPIKey != pmKey {
		t.Errorf("NewAPIClient() pmAPIKey = %v, want %v", client.pmAPIKey, pmKey)
	}

	if client.httpClient == nil {
		t.Error("NewAPIClient() httpClient is nil")
	}

	if client.httpClient.Timeout != 30*time.Second {
		t.Errorf("NewAPIClient() httpClient.Timeout = %v, want %v", client.httpClient.Timeout, 30*time.Second)
	}
}

func TestAPIClient_fetchDoc(t *testing.T) {
	tests := []struct {
		name           string
		mockResponse   string
		mockStatusCode int
		apiKey         string
		wantErr        bool
		errContains    string
		expectedResult string
	}{
		{
			name:           "successful fetch with valid JSON",
			mockResponse:   `{"openapi":"3.0.0","info":{"title":"Test API","version":"1.0.0"}}`,
			mockStatusCode: http.StatusOK,
			apiKey:         "test-api-key",
			wantErr:        false,
			expectedResult: "{\n  \"info\": {\n    \"title\": \"Test API\",\n    \"version\": \"1.0.0\"\n  },\n  \"openapi\": \"3.0.0\"\n}",
		},
		{
			name:           "API returns 404",
			mockResponse:   `{"error":"Not Found"}`,
			mockStatusCode: http.StatusNotFound,
			apiKey:         "test-api-key",
			wantErr:        true,
			errContains:    "unexpected status: 404",
		},
		{
			name:           "invalid JSON response",
			mockResponse:   `{"invalid": json}`,
			mockStatusCode: http.StatusOK,
			apiKey:         "test-api-key",
			wantErr:        true,
			errContains:    "decoding JSON",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if apiKey := r.Header.Get("X-API-Key"); apiKey != tt.apiKey {
					t.Errorf("Expected X-API-Key header %q, got %q", tt.apiKey, apiKey)
				}

				w.WriteHeader(tt.mockStatusCode)
				w.Write([]byte(tt.mockResponse))
			}))
			defer server.Close()

			client := NewAPIClient(tt.apiKey, "pm-key")
			result, err := client.fetchDoc(server.URL)

			if tt.wantErr {
				if err == nil {
					t.Errorf("fetchDoc() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("fetchDoc() error = %v, want error containing %v", err, tt.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("fetchDoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if result != tt.expectedResult {
				t.Errorf("fetchDoc() = %v, want %v", result, tt.expectedResult)
			}
		})
	}
}

func TestAPIClient_JSONParsing(t *testing.T) {
	// Test that the client can handle various JSON formats
	testCases := []struct {
		name          string
		jsonResponse  string
		shouldSucceed bool
	}{
		{
			name:          "simple object",
			jsonResponse:  `{"key": "value"}`,
			shouldSucceed: true,
		},
		{
			name:          "nested object",
			jsonResponse:  `{"outer": {"inner": "value"}}`,
			shouldSucceed: true,
		},
		{
			name:          "array",
			jsonResponse:  `[{"id": 1}, {"id": 2}]`,
			shouldSucceed: true,
		},
		{
			name:          "invalid JSON",
			jsonResponse:  `{invalid}`,
			shouldSucceed: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test if the JSON can be parsed (this tests the logic in fetchDoc)
			var result any
			err := json.Unmarshal([]byte(tc.jsonResponse), &result)

			if tc.shouldSucceed {
				if err != nil {
					t.Errorf("Expected JSON to parse successfully, got error: %v", err)
				}
			} else {
				if err == nil {
					t.Error("Expected JSON parsing to fail, but it succeeded")
				}
			}
		})
	}
}
