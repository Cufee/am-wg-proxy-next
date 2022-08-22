package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"aftermath.link/repo/am-wg-proxy/logs"
	_ "github.com/joho/godotenv/autoload"
)

func HttpRequest(url, method string, proxy *url.URL, headers map[string]string, payload []byte, target interface{}) (int, error) {
	var err error
	var bodyBytes []byte
	var resp *http.Response
	defer func() {
		// Logging
		if err != nil || resp == nil || resp.StatusCode != http.StatusOK {
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
	transport := &http.Transport{}
	if proxy != nil {
		transport.Proxy = http.ProxyURL(proxy)
	}

	client := &http.Client{
		// Timeout:   5 * time.Second,
		Transport: transport,
	}
	resp, err = client.Do(req)
	if resp == nil {
		logs.Error(logs.Wrap(err, "client.Do failed").Error())
		return 0, errors.New("client.Do returned nil response")
	}
	if err != nil {
		return resp.StatusCode, logs.Wrap(err, "client.Do failed")
	}
	// Read body
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, logs.Wrap(err, "ioutil.ReadAll failed")
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
