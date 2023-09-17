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

func UNKNOWN_METHOD(method string) Method {
	return Method(method)
}

type TLSVersion string

var (
	TLSV1_1 TLSVersion = "tlsv1.1"
	TLSV1_2 TLSVersion = "tlsv1.2"
	TLSV1_3 TLSVersion = "tlsv1.3"
)

type HTTPVersion string

var (
	HTTPVersion1_0 HTTPVersion = "http1.0"
	HTTPVersion1_1 HTTPVersion = "http1.1"
	HTTPVersion2_0 HTTPVersion = "http2.0"
	HTTPVersion3_0 HTTPVersion = "http3.0"
)

type RequestConfig struct {
	Url     string
	Method  Method
	Headers map[string]string
	Data    string
	Http    HTTPVersion
	Tls     TLSVersion
}

type OutputConfig struct {
	IsHead bool
}

type Config struct {
	request RequestConfig
	output  OutputConfig
}

func NewConfig(request RequestConfig, output OutputConfig) Config {
	return Config{
		request: request,
		output:  output,
	}
}
