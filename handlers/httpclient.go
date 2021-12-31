package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"aftermath.link/repo/logs"
)

func HTTPRequest(url string, method string, headers map[string]string, payload []byte, target interface{}) (int, error) {
	var err error
	var bodyBytes []byte
	var resp *http.Response
	defer func() {
		// Logging
		if err != nil || resp.StatusCode != http.StatusOK {
			logs.Warning("URL: %v", url)
			logs.Warning("Headers: %v", headers)
			logs.Warning("Payload: %v", string(payload))
			logs.Warning("Response: %v", string(bodyBytes))
			logs.Warning("Error: %v", err)
		}
	}()

	// Prep request
	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer(payload))
	if err != nil {
		return 0, logs.Wrap(err, "http.NewRequest failed")
	}

	// Set headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Set payload headers
	if payload != nil {
		req.Header.Set("content-type", "application/json")
	}

	// Send request
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return resp.StatusCode, logs.Wrap(err, "client.Do failed")
	}
	// Read body
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, logs.Wrap(err, "ioutil.ReadAll failed")
	}

	// Check response code
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, fmt.Errorf("HTTP status code: %v", resp.StatusCode)
	}

	// Decode
	if target != nil {
		err = json.Unmarshal(bodyBytes, target)
		if err != nil {
			return resp.StatusCode, logs.Wrap(err, "json.Unmarshal failed")
		}
	}
	return resp.StatusCode, nil
}
