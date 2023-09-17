package core

import "fmt"

func Run(config Config) error {
	client, err := NewClient(config.request)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	response, headers, err := client.Do()
	if err != nil {
		return fmt.Errorf("failed to request: %w", err)
	}

	output := NewOutput(config.output)
	err = output.Do(response, headers)
	if err != nil {
		return fmt.Errorf("failed to output: %w", err)
	}

	return nil
}
