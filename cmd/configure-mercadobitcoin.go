package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"fmt"
)

var tapi_id string
var tapi_secret string

func init() {
	configureCmd.AddCommand(mercadobitcoinCmd)
	mercadobitcoinCmd.Flags().StringVarP(&tapi_id, "tapi-id", "i", "", "TAPI_ID of your account")
	mercadobitcoinCmd.Flags().StringVarP(&tapi_secret, "tapi-secret", "s", "", "TAPI_SECRET of your account")
}

type config struct {
    TapiId      string
    TapiSecret  string
}

var C config

func SaveConfig(args []string) {

	if tapi_id != "" {
		viper.Set("tapi_id", tapi_id)
	}
	if tapi_secret != "" {
		viper.Set("tapi_secret", tapi_secret)
	}
	viper.WriteConfig()

	fmt.Println("Saved mercadobitcoin configuration to config.properties")
}

var mercadobitcoinCmd = &cobra.Command {
	Use: 	 "mercadobitcoin",
	Short: 	 "Configure Mercadobitcoin TAPI credentials",
	Run: func(cmd *cobra.Command, args []string) {
		SaveConfig(args)
	},
}
