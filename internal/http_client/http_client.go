package http_client

import (
	"base/types"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpClient struct{}

func BuildHttpClient() *HttpClient {
	return &HttpClient{}
}

func (client *HttpClient) Get(url string) ([]byte, *http.Response, error) {
	return client.MakeRequest("GET", url, types.RequestOpts{
		Body:    nil,
		Headers: nil,
	})
}
func (client *HttpClient) Post(url string, body interface{}) ([]byte, *http.Response, error) {
	return client.MakeRequest("POST", url, types.RequestOpts{
		Body:    body,
		Headers: nil,
	})
}
func (client *HttpClient) Put(url string, body interface{}) ([]byte, *http.Response, error) {
	return client.MakeRequest("PUT", url, types.RequestOpts{
		Body:    body,
		Headers: nil,
	})

}
func (client *HttpClient) Delete(url string, body interface{}) ([]byte, *http.Response, error) {
	return client.MakeRequest("DELETE", url, types.RequestOpts{
		Body:    body,
		Headers: nil,
	})
}

func (client *HttpClient) MakeRequest(method string, url string, opts types.RequestOpts) ([]byte, *http.Response, error) {

	var buf bytes.Buffer
	c := &http.Client{}

	if opts.Body != nil {
		err := json.NewEncoder(&buf).Encode(opts.Body)
		if err != nil {
			return nil, nil, err
		}
	}
	req, err := http.NewRequest(method, url, &buf)
	if err != nil {
		return nil, nil, err
	}

	if opts.Headers != nil {
		for k, v := range opts.Headers {
			req.Header.Set(k, v[0])
			if len(v) > 1 {
				for idx, elem := range v {
					if idx != 0 {
						req.Header.Add(k, elem)
					}
				}
			}
		}
	}

	fmt.Println(req.Header.Clone())

	resp, err := c.Do(req)
	if err != nil {
		return nil, nil, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return bytes, resp, nil

}
