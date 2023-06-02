package helpers

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// PostHTTPRequest ...
func PostHTTPRequest(url, authorization string, data []byte) ([]byte, error) {
	// Prepare request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))

	// Add authorization header to the req
	req.Header.Add("Authorization", authorization)
	req.Header.Set("Content-Type", "application/json")

	// Send request using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return []byte(body), nil
}

// GetHTTPRequest ...
func GetHTTPRequest(url, authorization string) ([]byte, error) {
	// Prepare request body
	req, err := http.NewRequest("GET", url, nil)

	// Add authorization header to the req
	req.Header.Add("Authorization", authorization)
	req.Header.Set("Content-Type", "application/json")

	// Send request using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return []byte(body), nil
}
