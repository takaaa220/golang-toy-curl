package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/takaaa220/golang-toy-curl/config"
)

var (
	rootCmd = &cobra.Command{
		Use:                   "tcurl [SUBCOMMAND] [OPTIONS] [url]",
		DisableFlagsInUseLine: true,
		Short:                 "A tool to make requests.",
	}

	methodStr string
	headers   []string
	data      string
	isHead    bool
	tls1_1    bool
	tls1_2    bool
	tls1_3    bool
	http1_0   bool
	http1_1   bool
	http2_0   bool
	http3_0   bool
)

type Run func(config config.Config) error

func wrap(run Run) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("url is required")
		}

		url := args[0]
		if url == "" || !strings.HasPrefix(url, "http") {
			return fmt.Errorf("invalid url: %s", url)
		}

		headers, err := parseHeaders(headers)
		if err != nil {
			return err
		}

		method := config.NetMethod(methodStr)
		http := getHttp(http1_0, http1_1, http2_0, http3_0)
		tls := getTls(tls1_1, tls1_2, tls1_3)

		return run(config.NewConfig(
			config.RequestConfig{
				Url:     url,
				Method:  method,
				Headers: headers,
				Data:    data,
				Http:    http,
				Tls:     tls,
			},
			config.OutputConfig{
				IsHead: isHead,
			},
		))
	}
}

func Execute(runCurl Run, runToGo Run) error {
	rootCmd.RunE = wrap(runCurl)
	toGoCmd.RunE = wrap(runToGo)

	return rootCmd.Execute()
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
	// request option
	rootCmd.PersistentFlags().StringVarP(&methodStr, "request", "X", "GET", "Specify request method")
	rootCmd.PersistentFlags().StringSliceVarP(&headers, "header", "H", []string{}, "Set request headers")
	rootCmd.PersistentFlags().StringVarP(&data, "data", "d", "", "HTTP POST data")

	// protocol option (tls)
	rootCmd.PersistentFlags().BoolVar(&tls1_1, "tlsv1.1", false, "Use TLSv1.1 or greater")
	rootCmd.PersistentFlags().BoolVar(&tls1_2, "tlsv1.2", false, "Use TLSv1.2 or greater")
	rootCmd.PersistentFlags().BoolVar(&tls1_3, "tlsv1.3", true, "Use TLSv1.3 or greater")

	// protocol option (http)
	rootCmd.PersistentFlags().BoolVar(&http1_0, "http1.0", false, "Use HTTP/1.0")
	rootCmd.PersistentFlags().BoolVar(&http1_1, "http1.1", false, "Use HTTP/1.1")
	rootCmd.PersistentFlags().BoolVar(&http2_0, "http2.0", true, "Use HTTP/2.0")
	rootCmd.PersistentFlags().BoolVar(&http3_0, "http3.0", false, "Use HTTP/3.0")

	// output option
	rootCmd.PersistentFlags().BoolVarP(&isHead, "head", "I", false, "Show response headers only")

	rootCmd.AddCommand(toGoCmd)
}

func getTls(v1_1, v1_2, v1_3 bool) config.TLSVersion {
	if v1_2 {
		return config.TLSV1_2
	}
	if v1_1 {
		return config.TLSV1_1
	}

	return config.TLSV1_3
}

func getHttp(v1, v1_1, v2, v3 bool) config.HTTPVersion {
	if v3 {
		return config.HTTPV3_0
	}
	if v1_1 {
		return config.HTTPV1_1
	}
	if v1 {
		return config.HTTPV1_0
	}

	return config.HTTPV2_0
}
