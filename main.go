package main

import (
	"os"

	"github.com/takaaa220/golang-toy-curl/cmd"
	"github.com/takaaa220/golang-toy-curl/core"
)

func main() {
	err := cmd.Execute(core.Run)
	if err != nil {
		os.Exit(1)
	}
}
