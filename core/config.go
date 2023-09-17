package core

type Method string

var (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	PATCH   Method = "PATCH"
	DELETE  Method = "DELETE"
	OPTIONS Method = "OPTIONS"
	QUERY   Method = "QUERY"
)

type RequestConfig struct {
	Url     string
	Method  Method
	Headers map[string]string
	Data    string
}

func newRequestConfig(url string, method Method, headers map[string]string, data string) RequestConfig {
	return RequestConfig{
		Url:     url,
		Method:  method,
		Headers: headers,
		Data:    data,
	}
}

type OutputConfig struct {
	IsHead bool
}

func newOutputConfig(isHead bool) OutputConfig {
	return OutputConfig{
		IsHead: isHead,
	}
}

type Config struct {
	request RequestConfig
	output  OutputConfig
}

func NewConfig(url string, method Method, headers map[string]string, data string, isHead bool) Config {
	return Config{
		request: newRequestConfig(url, method, headers, data),
		output:  newOutputConfig(isHead),
	}
}
