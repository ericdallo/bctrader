package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(configureCmd)
}

var configureCmd = &cobra.Command {
	Use: 	 "configure",
	Short: 	 "Configure credentials for APIs http calls",
	Aliases: []string { "configuration", "config", "cfg"},
}
