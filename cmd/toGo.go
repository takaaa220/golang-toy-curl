package cmd

import (
	"github.com/spf13/cobra"
)

var toGoCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert curl command into Go code.",
}
