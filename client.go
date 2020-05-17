package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// APIClient definition
type APIClient struct {
	URL     string
	Headers map[string]string
	Client  *http.Client
}

// New - create API client instance
func New(url string, headers map[string]string) *APIClient {
	return &APIClient{
		URL:     url,
		Headers: headers,
		Client:  &http.Client{},
	}
}

// BuildRequest - returns a prepared HTTP request on success
func (ac *APIClient) BuildRequest(method string, body interface{}) (apiReq *http.Request, err error) {
	// JSON encode request body
	jsonReqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Initialize the request
	apiReq, err = http.NewRequest(method, ac.URL, bytes.NewBuffer(jsonReqBody))
	if err != nil {
		return nil, err
	}

	// Set request headers
	for hKey, hVal := range ac.Headers {
		apiReq.Header.Set(hKey, hVal)
	}

	return apiReq, nil
}

// Do - execute the request
func (ac *APIClient) Do(apiReq *http.Request, respBody interface{}) (apiResp *http.Response, err error) {
	// Perform the request
	apiResp, err = ac.Client.Do(apiReq)
	if err != nil {
		return nil, err
	}
	// Close the response body on function return
	defer apiResp.Body.Close()

	// Read the response
	bResponse, err := ioutil.ReadAll(apiResp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal response
	if respBody != nil && len(bResponse) != 0 {
		err = json.Unmarshal(bResponse, respBody)
		if err != nil {
			return nil, err
		}
	}

	return apiResp, nil
}
