# Simple REST API client

## Description
Lightweight REST API client using core packages only

## Installation
```
$ go get github.com/andreimilas/rest-api-client
```

## Usage
#### GET
```
import (
    apiClient "github.com/andreimilas/rest-api-client"
)

...
apiURL := "<API_URL>"

// Build request headers
rHeaders := map[string]string{
    "Content-Type": "application/json",
    ...
}

// Build the API request
client := apiClient.New(apiURL, rHeaders)
apiRequest, err := client.BuildRequest("GET", nil)
if err != nil {
    return nil, err
}

// Send the API request
responseBody = &map[string]interface{}{}
apiResponse, err := client.Do(apiRequest, responseBody)
if err != nil {
    return nil, err
}
```

#### POST
```
import (
    apiClient "github.com/andreimilas/rest-api-client"
)

...
apiURL := "<API_URL>"

// Build request headers
rHeaders := map[string]string{
    "Content-Type": "application/json",
    ...
}

// Build request body
rBody := &map[string]interface{}{"data": ...}

// Build the API request
client := apiClient.NewAPIClient(apiURL, rHeaders)
apiRequest, err := client.BuildRequest("POST", rBody)
if err != nil {
    return nil, err
}

// Send the API request
responseBody = &map[string]interface{}{}
apiResponse, err := client.Do(apiRequest, responseBody)
if err != nil {
    return nil, err
}
```

#### PUT
```
import (
    apiClient "github.com/andreimilas/rest-api-client"
)

...
apiURL := "<API_URL>"

// Build request headers
rHeaders := map[string]string{
    "Content-Type": "application/json",
    ...
}

// Build request body
rBody := &map[string]interface{}{"data": ...}

// Build the API request
client := apiClient.NewAPIClient(apiURL, rHeaders)
apiRequest, err := client.BuildRequest("PUT", rBody)
if err != nil {
    return nil, err
}

// Send the API request
responseBody = &map[string]interface{}{}
apiResponse, err := client.Do(apiRequest, responseBody)
if err != nil {
    return nil, err
}
```

#### DELETE
```
import (
    apiClient "github.com/andreimilas/rest-api-client"
)

...
apiURL := "<API_URL>"

// Build request headers
rHeaders := map[string]string{
    "Content-Type": "application/json",
    ...
}

// Build the API request
client := apiClient.New(apiURL, rHeaders)
apiRequest, err := client.BuildRequest("DELETE", nil)
if err != nil {
    return nil, err
}

// Send the API request
responseBody = &map[string]interface{}{}
apiResponse, err := client.Do(apiRequest, responseBody)
if err != nil {
    return nil, err
}
```