package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command {
	Use: 	"version",
	Short: 	"Print the version of BCtrader",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(rootCmd.Use + " current version: " + VERSION + "\n") 
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}