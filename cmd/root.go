package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/takaaa220/golang-toy-curl/core"
)

var (
	rootCmd = &cobra.Command{
		Use:                   "tcurl [OPTIONS] [url]",
		DisableFlagsInUseLine: true,
		Short:                 "A tool to make requests.",
	}

	method  string
	headers []string
	data    string
)

type Run func(config core.Config) error

func Execute(run Run) error {
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("url is required")
		}

		url := args[0]
		if url == "" || !strings.HasPrefix(url, "http") {
			return fmt.Errorf("invalid url: %s", url)
		}

		method, err := parseMethod(method)
		if err != nil {
			return err
		}

		headers, err := parseHeaders(headers)
		if err != nil {
			return err
		}

		return run(core.NewConfig(url, method, headers, data))
	}

	return rootCmd.Execute()
}

func parseMethod(method string) (core.Method, error) {
	switch method {
	case "GET":
		return core.GET, nil
	case "POST":
		return core.POST, nil
	case "PUT":
		return core.PUT, nil
	case "PATCH":
		return core.PATCH, nil
	case "DELETE":
		return core.DELETE, nil
	case "OPTIONS":
		return core.OPTIONS, nil
	case "QUERY":
		return core.QUERY, nil
	default:
		return "", fmt.Errorf("invalid method: %s", method)
	}
}

func parseHeaders(headers []string) (map[string]string, error) {
	ma := make(map[string]string)

	for _, header := range headers {
		name, value, err := splitHeader(header)
		if err != nil {
			return nil, err
		}

		ma[name] = value
	}

	return ma, nil
}

func splitHeader(header string) (name string, value string, err error) {
	split := strings.Split(header, ":")
	if len(split) != 2 {
		fmt.Println("Invalid header format.")
		return "", "", fmt.Errorf("invalid header: %s", header)
	}

	return split[0], split[1], nil
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&method, "request", "X", "GET", "Specify request method")
	rootCmd.PersistentFlags().StringSliceVarP(&headers, "header", "H", []string{}, "Set request headers")
	rootCmd.PersistentFlags().StringVarP(&data, "data", "d", "", "HTTP POST data")
}
