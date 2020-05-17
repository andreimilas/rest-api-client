package client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	// Multiple test cases
	var tests = []struct {
		name              string
		url               string
		headers           map[string]string
		expectedAPIClient *APIClient
	}{
		{"With URL & Headers",
			"http://localhost",
			map[string]string{"Content-Type": "application/json", "Cache-Control": "no-cache"},
			&APIClient{URL: "http://localhost", Headers: map[string]string{"Content-Type": "application/json", "Cache-Control": "no-cache"}},
		},
		{"With URL & Header",
			"http://localhost",
			map[string]string{"Content-Type": "application/json"},
			&APIClient{URL: "http://localhost", Headers: map[string]string{"Content-Type": "application/json"}},
		},
		{"With URL",
			"http://localhost",
			map[string]string{},
			&APIClient{URL: "http://localhost"},
		},
	}
	for _, test := range tests {
		t.Run("New API Client", func(t *testing.T) {
			// Initialize API client
			apiClient := New(test.url, test.headers)
			if apiClient.URL != test.expectedAPIClient.URL ||
				(!reflect.DeepEqual(apiClient.Headers, test.expectedAPIClient.Headers) &&
					!(len(apiClient.Headers) == 0 && len(test.expectedAPIClient.Headers) == 0)) {
				t.Errorf("Received: %v; Expected: %v.", apiClient, test.expectedAPIClient)
			}
		})
	}
}

func TestBuildRequest(t *testing.T) {
	// Multiple test cases
	var tests = []struct {
		name            string
		method          string
		body            interface{}
		apiClient       *APIClient
		expectedRequest *http.Request
	}{
		{"POST",
			"POST",
			map[string]interface{}{"a": "1", "b": 2, "c": true, "d": []string{"aa", "bb", "cc"}},
			New("http://localhost", map[string]string{}),
			mockRequest("POST", "http://localhost", map[string]interface{}{"a": "1", "b": 2, "c": true, "d": []string{"aa", "bb", "cc"}}),
		},
		{"GET",
			"GET",
			nil,
			New("http://localhost", map[string]string{}),
			mockRequest("GET", "http://localhost", nil),
		},
		{"PUT",
			"PUT",
			nil,
			New("http://localhost", map[string]string{}),
			mockRequest("PUT", "http://localhost", nil),
		},
		{"PATCH",
			"PATCH",
			nil,
			New("http://localhost", map[string]string{}),
			mockRequest("PATCH", "http://localhost", nil),
		},
		{"DELETE",
			"DELETE",
			nil,
			New("http://localhost", map[string]string{}),
			mockRequest("DELETE", "http://localhost", nil),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Build the request and check corresponding attributes
			req, _ := test.apiClient.BuildRequest(test.method, test.body)
			if req.Method != test.expectedRequest.Method ||
				!reflect.DeepEqual(req.URL, test.expectedRequest.URL) ||
				!reflect.DeepEqual(req.Body, test.expectedRequest.Body) {
				t.Errorf("Received: %v; Expected: %v", req, test.expectedRequest)
			}
		})
	}
}

func TestDo(t *testing.T) {
	t.Run("Do Request", func(t *testing.T) {
		// Start an HTTP server for testing purposes
		go func() {
			http.HandleFunc("/v1/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
			log.Fatal(http.ListenAndServe(":8080", nil))
		}()

		// Init API client
		apiClient := New("http://localhost:8080/v1/health", map[string]string{"Content-Type": "application/json", "Cache-Control": "no-cache"})
		// Build and send GET request
		req, _ := apiClient.BuildRequest("GET", nil)
		resp, err := apiClient.Do(req, nil)
		if err != nil {
			t.Errorf("Received Error: %s; Expected: ''", err.Error())
		} else if resp.StatusCode != 200 {
			t.Errorf("Received Status Code: %d; Expected: %d", resp.StatusCode, 200)
		}
	})
}

// Helper function which creates a request instance using the given arguments
func mockRequest(method string, url string, body interface{}) *http.Request {
	jsonReqBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonReqBody))
	return req
}
