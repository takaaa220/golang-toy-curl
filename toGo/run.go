package toGo

import (
	"fmt"
	"strings"

	"github.com/takaaa220/golang-toy-curl/config"
)

func convertMethod(method config.Method) string {
	return string(method)
}

func convertHeaders(headers map[string]string, indent int) string {
	headersStr := "map[string]string{\n"
	for k, v := range headers {
		headersStr += fmt.Sprintf("\t\"%s\": \"%s\",\n", k, v)
	}
	headersStr += "}"

	return strings.ReplaceAll(headersStr, "\n", "\n\t")
}

func convertTLSVersion(t config.TLSVersion) string {
	switch t {
	case config.TLSV1_1:
		return "tls.VersionTLS11"
	case config.TLSV1_2:
		return "tls.VersionTLS12"
	case config.TLSV1_3:
		return "tls.VersionTLS13"
	default:
		panic(fmt.Sprintf("unknown tls version: %s", t))
	}
}

func Run(cfg config.Config) error {
	code := template(
		templateArgs{
			method:     convertMethod(cfg.Request.Method),
			url:        cfg.Request.Url,
			data:       cfg.Request.Data,
			headers:    convertHeaders(cfg.Request.Headers, 1),
			tlsVersion: convertTLSVersion(cfg.Request.Tls),
		},
	)

	output(code, cfg.Output)

	return nil
}
