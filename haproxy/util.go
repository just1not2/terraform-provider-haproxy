// Copyright: (c) 2022, Justin Béra (@just1not2) <me@just1not2.org>
// Mozilla Public License Version 2.0 (see LICENSE or https://www.mozilla.org/en-US/MPL/2.0/)

package haproxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type HAProxyAuthentication struct {
	username string
	password string
}

type HAProxyClient struct {
	auth   *HAProxyAuthentication
	client *http.Client
	mutex  *sync.Mutex
	url    string
}

func NewHAProxyClient(url, username, password interface{}) *HAProxyClient {
	var auth *HAProxyAuthentication

	// Sets authentication if it is defined in the provider options
	if username != nil && password != nil {
		auth = &HAProxyAuthentication{
			username: username.(string),
			password: password.(string),
		}
	}

	// Creates the HAProxy client
	client := &HAProxyClient{
		url:    url.(string),
		client: &http.Client{Timeout: 10 * time.Second},
		auth:   auth,
		mutex:  &sync.Mutex{},
	}

	return client
}

func (client *HAProxyClient) Request(method string, uri string, body *map[string]string, returnBody *map[string]interface{}) error {
	var version string

	// Gets HAProxy API configuration version for stateful requests
	if method != "GET" {
		var versionReturnBody map[string]interface{}

		// Uses mutexes to avoid conflicts in configuration versions
		client.mutex.Lock()
		defer client.mutex.Unlock()

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/services/haproxy/configuration/raw", client.url), nil)
		if err != nil {
			return err
		}

		// Sends the request to get the configuration version
		if err = client.SendRequest(req, &versionReturnBody); err != nil {
			return err
		}

		version = fmt.Sprintf("?version=%v", versionReturnBody["_version"].(float64))
	}

	// Sets the body into the buffer
	sender := &bytes.Buffer{}
	if body != nil {
		jsonData, _ := json.Marshal(body)
		sender = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s%s", client.url, uri, version), sender)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	return client.SendRequest(req, returnBody)
}

func (client *HAProxyClient) SendRequest(req *http.Request, returnBody *map[string]interface{}) error {
	// Sets authentication configuration
	if client.auth != nil {
		req.SetBasicAuth(client.auth.username, client.auth.password)
	}

	// Sends the HTTP request
	res, err := client.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Sets the body into a JSON
	json.NewDecoder(res.Body).Decode(returnBody)

	// Throws an error if the HAProxy returns an error message
	if (*returnBody)["code"] != nil {
		return fmt.Errorf("error %v: %s", (*returnBody)["code"].(float64), (*returnBody)["message"].(string))
	}

	return nil
}
