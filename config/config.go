package config

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
	IsHead  bool
	Http    HTTPVersion
	Tls     TLSVersion
}

type OutputConfig struct {
}

type Config struct {
	Request RequestConfig
	Output  OutputConfig
}

func NewConfig(request RequestConfig, output OutputConfig) Config {
	return Config{
		Request: request,
		Output:  output,
	}
}
