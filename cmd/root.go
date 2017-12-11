package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"log"
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
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Panic(err)
	}

	VERSION = version

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}