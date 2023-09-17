package curl

import (
	"fmt"

	"github.com/takaaa220/golang-toy-curl/config"
)

type Output struct {
	config config.OutputConfig
}

func NewOutput(config config.OutputConfig) Output {
	return Output{
		config: config,
	}
}

func (o Output) Do(body string, headers map[string]string) error {
	if o.config.IsHead {
		return o.head(headers)
	}

	return o.body(body)
}

func (o Output) head(headers map[string]string) error {
	for k, v := range headers {
		fmt.Printf("%s: %s\n", k, v)
	}

	return nil
}

func (o Output) body(body string) error {
	fmt.Println(body)

	return nil
}
