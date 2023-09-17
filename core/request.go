package core

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	request *http.Request
}

func NewClient(config RequestConfig) (*Client, error) {
	req, err := http.NewRequest(string(config.Method), config.Url, bytes.NewBuffer([]byte(config.Data)))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range config.Headers {
		req.Header.Set(k, v)
	}

	return &Client{
		request: req,
	}, nil
}

func (r Client) Do() (string, map[string]string, error) {
	client := &http.Client{}

	resp, err := client.Do(r.request)
	if err != nil {
		return "", map[string]string{}, fmt.Errorf("failed to request: %w", err)
	}
	defer resp.Body.Close()

	headers, err := makeHeaders(resp), nil
	if err != nil {
		return "", map[string]string{}, err
	}

	body, err := makeBody(resp)
	if err != nil {
		return "", map[string]string{}, err
	}

	return body, headers, nil
}

func makeBody(response *http.Response) (string, error) {
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(bodyBytes), nil
}

func makeHeaders(response *http.Response) map[string]string {
	ma := make(map[string]string)

	for k, v := range response.Header {
		ma[k] = v[0]
	}

	return ma
}
