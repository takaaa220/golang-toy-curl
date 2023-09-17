package main

import (
	"os"

	"github.com/takaaa220/golang-toy-curl/cmd"
	"github.com/takaaa220/golang-toy-curl/curl"
	"github.com/takaaa220/golang-toy-curl/toGo"
)

func main() {
	err := cmd.Execute(curl.Run, toGo.Run)
	if err != nil {
		os.Exit(1)
	}
}
