package core

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Url     string
	Method  Method
	Headers map[string]string
	Data    string
}

func Do(config Config) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(string(config.Method), config.Url, bytes.NewBuffer([]byte(config.Data)))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range config.Headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(bodyBytes), nil
}
