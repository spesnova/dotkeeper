/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dotkeeper",
	Short: "A declarative dotfiles management tool",
	Long: `dotkeeper is a command-line tool for managing your dotfiles declaratively.

It creates and manages symbolic links based on the configuration defined in your
dotfiles.yaml file. The tool ensures idempotent operations, making it safe to run
multiple times.

Example configuration (dotfiles.yaml):
    symlinks:
      - src: bash/bashrc
        dst: ~/.bashrc
      - src: vim/vimrc
        dst: ~/.vimrc
      - src: vim/vim
        dst: ~/.vim
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(applyCmd)
}
