package http_client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type HttpClient struct{}

func BuildHttpClient() *HttpClient {

	return &HttpClient{}
}

func (client *HttpClient) Get(url string) ([]byte, error) {
	bytes, _ := makeRequest("GET", url, nil)
	return bytes, nil
}
func (client *HttpClient) Post(url string, body interface{}) ([]byte, error) {
	bytes, _ := makeRequest("POST", url, body)
	return bytes, nil
}
func (client *HttpClient) Put(url string, body interface{}) ([]byte, error) {
	bytes, _ := makeRequest("PUT", url, body)
	return bytes, nil
}
func (client *HttpClient) Delete(url string, body interface{}) (bool, error) {
	_, err := makeRequest("DELETE", url, body)
	return err == nil, nil
}

func makeRequest(method string, url string, body interface{}) ([]byte, error) {
	var buff bytes.Buffer
	client := &http.Client{}

	if body != nil {
		err := json.NewEncoder(&buff).Encode(body)
		if err != nil {
			return nil, nil
		}
	}
	req, err := http.NewRequest(method, url, &buff)
	if err != nil {
		return nil, nil
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}

	return bytes, nil

}
