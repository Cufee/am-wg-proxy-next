package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/byvko-dev/am-core/logs"
)

func Send(url string, method string, headers map[string]string, payload []byte, target interface{}) (int, error) {
	var err error
	var bodyBytes []byte
	var resp *http.Response
	defer func() {
		// Logging
		if err != nil || (resp != nil && resp.StatusCode != http.StatusOK) {
			logs.Error("URL: %v", url)
			logs.Error("Headers: %v", headers)
			logs.Error("Payload: %v", string(payload))
			logs.Error("Response: %v", string(bodyBytes))
			logs.Error("Error: %v", err)
		}
	}()

	// Prep request
	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer(payload))
	if err != nil {
		return 0, err
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
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		if resp != nil {
			return resp.StatusCode, err
		}
		return 0, err
	}
	// Read body
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}
	// Check response code
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, fmt.Errorf(resp.Status)
	}

	// Decode
	if target != nil {
		err = json.Unmarshal(bodyBytes, target)
	}
	return resp.StatusCode, err
}
