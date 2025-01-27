package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const VERSION = "0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the  CLI version",
	RunE:  runVersion,
}

func runVersion(cmd *cobra.Command, args []string) error {
	fmt.Println(VERSION)
	return nil
}
