package cmd

import (
	"fmt"

	"github.com/spesnova/dotkeeper/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the CLI version",
	RunE:  runVersion,
}

func runVersion(cmd *cobra.Command, args []string) error {
	fmt.Println(version.GetVersion())
	return nil
}
