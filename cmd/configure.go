package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(configureCmd)
}

func Configure(args []string) {
	
}

var configureCmd = &cobra.Command {
	Use: 	 "configure",
	Short: 	 "Configure credentials for APIs http calls",
	Aliases: []string { "config", "cfg"},
	Run: func(cmd *cobra.Command, args []string) {
		Configure(args)
	},
}
