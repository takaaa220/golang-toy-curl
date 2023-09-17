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

		headers, err := parseHeaders(headers)
		if err != nil {
			return err
		}

		method := core.NetMethod(methodStr)
		http := getHttp(http1_0, http1_1, http2_0, http3_0)
		tls := getTls(tls1_1, tls1_2, tls1_3)

		return run(core.NewConfig(
			core.RequestConfig{
				Url:     url,
				Method:  method,
				Headers: headers,
				Data:    data,
				Http:    http,
				Tls:     tls,
			},
			core.OutputConfig{
				IsHead: isHead,
			},
		))
	}

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
	rootCmd.Flags().StringVarP(&methodStr, "request", "X", "GET", "Specify request method")
	rootCmd.Flags().StringSliceVarP(&headers, "header", "H", []string{}, "Set request headers")
	rootCmd.Flags().StringVarP(&data, "data", "d", "", "HTTP POST data")

	// protocol option (tls)
	rootCmd.Flags().BoolVar(&tls1_1, "tlsv1.1", false, "Use TLSv1.1 or greater")
	rootCmd.Flags().BoolVar(&tls1_2, "tlsv1.2", false, "Use TLSv1.2 or greater")
	rootCmd.Flags().BoolVar(&tls1_3, "tlsv1.3", true, "Use TLSv1.3 or greater")

	// protocol option (http)
	rootCmd.Flags().BoolVar(&http1_0, "http1.0", false, "Use HTTP/1.0")
	rootCmd.Flags().BoolVar(&http1_1, "http1.1", false, "Use HTTP/1.1")
	rootCmd.Flags().BoolVar(&http2_0, "http2.0", true, "Use HTTP/2.0")
	rootCmd.Flags().BoolVar(&http3_0, "http3.0", false, "Use HTTP/3.0")

	// output option
	rootCmd.Flags().BoolVarP(&isHead, "head", "I", false, "Show response headers only")

}

func getTls(v1_1, v1_2, v1_3 bool) core.TLSVersion {
	if v1_2 {
		return core.TLSV1_2
	}
	if v1_1 {
		return core.TLSV1_1
	}

	return core.TLSV1_3
}

func getHttp(v1, v1_1, v2, v3 bool) core.HTTPVersion {
	if v3 {
		return core.HTTPV3_0
	}
	if v1_1 {
		return core.HTTPV1_1
	}
	if v1 {
		return core.HTTPV1_0
	}

	return core.HTTPV2_0
}
