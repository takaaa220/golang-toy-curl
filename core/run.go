package core

import "fmt"

func Run(config Config) error {
	response, err := Do(config)
	if err != nil {
		return fmt.Errorf("failed to request: %w", err)
	}

	output, err := Output(response, config)
	if err != nil {
		return fmt.Errorf("failed to output: %w", err)
	}

	fmt.Println(output)

	return nil
}
