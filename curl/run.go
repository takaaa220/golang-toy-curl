package curl

import (
	"fmt"

	"github.com/takaaa220/golang-toy-curl/config"
)

func Run(cfg config.Config) error {
	client, err := NewClient(cfg.Request)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	response, headers, err := client.Do()
	if err != nil {
		return fmt.Errorf("failed to request: %w", err)
	}

	output := NewOutput(cfg.Output)
	err = output.Do(response, headers)
	if err != nil {
		return fmt.Errorf("failed to output: %w", err)
	}

	return nil
}
