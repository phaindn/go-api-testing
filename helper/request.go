package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

func makeRequest(method string, url string, data *map[string]interface{}, token string, params *[]Parameter) (*http.Response, error) {

	client := http.Client{}

	var buffer *bytes.Buffer = nil

	// fmt.Println("data:", data, "buffer:", buffer)
	if data != nil {
		body, err := json.Marshal(*data)
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewBuffer(body)
	}

	endpoint := url

	// check parameters
	if params != nil {
		for _, p := range *params {
			endpoint = strings.Replace(endpoint, ":"+p.Key, p.Value, 1)
		}
	}

	req, err := http.NewRequest(method, endpoint, buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get method
func Get(url string, data *map[string]interface{}, token string) (*http.Response, error) {
	return makeRequest("GET", url, data, token, nil)
}

// Post method
func Post(url string, data *map[string]interface{}, token string) (*http.Response, error) {
	return makeRequest("POST", url, data, token, nil)
}

// Put method
func Put(url string, data *map[string]interface{}, token string) (*http.Response, error) {
	return makeRequest("PUT", url, data, token, nil)
}

// Delete method
func Delete(url string, data *map[string]interface{}, token string) (*http.Response, error) {
	return makeRequest("DELETE", url, data, token, nil)
}

// ParamGet method with params
func ParamGet(url string, data *map[string]interface{}, token string, params []Parameter) (*http.Response, error) {
	return makeRequest("GET", url, data, token, &params)
}

// ParamPost method with params
func ParamPost(url string, data *map[string]interface{}, token string, params []Parameter) (*http.Response, error) {
	return makeRequest("POST", url, data, token, &params)
}

// ParamPut method with params
func ParamPut(url string, data *map[string]interface{}, token string, params []Parameter) (*http.Response, error) {
	return makeRequest("PUT", url, data, token, &params)
}

// ParamDelete method with params
func ParamDelete(url string, data *map[string]interface{}, token string, params []Parameter) (*http.Response, error) {
	return makeRequest("DELETE", url, data, token, &params)
}
