package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	VERSION string
)

var rootCmd = &cobra.Command {
	Use: "bctrader",
	Short: "BCtrader is a bitcoin bot trader",
	Long: `BCtrader is a works buying and selling Bitcoins
		   winning and worthing it with the currency
		   appreciation
			`,
}

func Execute(version string) {
	VERSION = version

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}