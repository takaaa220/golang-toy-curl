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

type Config struct {
	Url     string
	Method  Method
	Headers map[string]string
	Data    string
}

func NewConfig(url string, method Method, headers map[string]string, data string) Config {
	return Config{
		Url:     url,
		Method:  method,
		Headers: headers,
		Data:    data,
	}
}
