package core

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

var (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	PATCH   Method = "PATCH"
	DELETE  Method = "DELETE"
	OPTIONS Method = "OPTIONS"
	QUERY   Method = "QUERY"
)

func UNKNOWN_METHOD(method string) Method {
	return Method(method)
}

type Method string

func NetMethod(method string) Method {
	switch method {
	case "GET":
		return GET
	case "POST":
		return POST
	case "PUT":
		return PUT
	case "PATCH":
		return PATCH
	case "DELETE":
		return DELETE
	case "OPTIONS":
		return OPTIONS
	case "QUERY":
		return QUERY
	default:
		return UNKNOWN_METHOD(method)
	}
}

type TLSVersion string

var (
	TLSV1_1 TLSVersion = "tlsv1.1"
	TLSV1_2 TLSVersion = "tlsv1.2"
	TLSV1_3 TLSVersion = "tlsv1.3"
)

type HTTPVersion string

var (
	HTTPV1_0 HTTPVersion = "http1.0"
	HTTPV1_1 HTTPVersion = "http1.1"
	HTTPV2_0 HTTPVersion = "http2.0"
	HTTPV3_0 HTTPVersion = "http3.0"
)

type RequestConfig struct {
	Url     string
	Method  Method
	Headers map[string]string
	Data    string
	Http    HTTPVersion
	Tls     TLSVersion
}

type Client struct {
	request   *http.Request
	transport *http.Transport
}

func NewClient(config RequestConfig) (*Client, error) {
	req, err := http.NewRequest(
		string(config.Method),
		config.Url,
		bytes.NewBuffer([]byte(config.Data)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range config.Headers {
		req.Header.Set(k, v)
	}

	// TODO: support http3
	if config.Http == HTTPV3_0 {
		return nil, fmt.Errorf("HTTP3 is not supported yet")
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: config.Tls.convert(),
		},
		ForceAttemptHTTP2: config.Http == HTTPV2_0,
	}

	return &Client{
		request:   req,
		transport: transport,
	}, nil
}

func (t TLSVersion) convert() uint16 {
	switch t {
	case TLSV1_1:
		return tls.VersionTLS11
	case TLSV1_2:
		return tls.VersionTLS12
	case TLSV1_3:
		return tls.VersionTLS13
	default:
		panic(fmt.Sprintf("unknown tls version: %s", t))
	}
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
