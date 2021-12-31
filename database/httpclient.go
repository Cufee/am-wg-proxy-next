package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"aftermath.link/repo/logs"
)

func httpRequest(url string, method string, request apiRequest, target interface{}) error {
	var err error
	var bodyBytes []byte
	var resp *http.Response

	payload, err := json.Marshal(request)
	if err != nil {
		return err
	}

	defer func() {
		// Logging
		fn := logs.Debug
		if err != nil || resp.StatusCode != http.StatusOK {
			fn = logs.Error
		}
		fn("URL: %v", url)
		fn("Payload: %v", string(payload))
		fn("Response: %v", string(bodyBytes))
		fn("Error: %v", err)
	}()

	// Prep request
	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-api-key", accessKey)

	// Send request
	client := &http.Client{}
	resp, err = client.Do(req)
	if resp == nil {
		return fmt.Errorf("no response received")
	}
	if err != nil {
		return err
	}
	// Read body
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// Check response code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%v", string(bodyBytes))
	}

	// Decode
	if target != nil {
		err = json.Unmarshal(bodyBytes, target)
	}
	return err
}
