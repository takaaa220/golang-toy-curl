package curl

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	"github.com/takaaa220/golang-toy-curl/config"
)

type Client struct {
	request   *http.Request
	transport *http.Transport
}

func NewClient(cfg config.RequestConfig) (*Client, error) {
	req, err := http.NewRequest(
		string(cfg.Method),
		cfg.Url,
		bytes.NewBuffer([]byte(cfg.Data)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range cfg.Headers {
		req.Header.Set(k, v)
	}

	// TODO: support http3
	if cfg.Http == config.HTTPV3_0 {
		return nil, fmt.Errorf("HTTP3 is not supported yet")
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: convertTLSVersion(cfg.Tls),
		},
		ForceAttemptHTTP2: cfg.Http == config.HTTPV2_0,
	}

	return &Client{
		request:   req,
		transport: transport,
	}, nil
}

func convertTLSVersion(t config.TLSVersion) uint16 {
	switch t {
	case config.TLSV1_1:
		return tls.VersionTLS11
	case config.TLSV1_2:
		return tls.VersionTLS12
	case config.TLSV1_3:
		return tls.VersionTLS13
	default:
		panic(fmt.Sprintf("unknown tls version: %s", t))
	}
}

func (r Client) Do() (string, map[string]string, error) {
	client := &http.Client{
		Transport: r.transport,
	}

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
