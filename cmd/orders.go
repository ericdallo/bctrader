package cmd

import (
	"github.com/spf13/cobra"
	"github.com/bctrader/api"
)

func init() {
	rootCmd.AddCommand(ordersCmd)
}

func ListOrders(args[] string) {
	api.GetOrder()
}

var ordersCmd = &cobra.Command {
	Use: 	 "orders",
	Short: 	 "List all your orders",
	Aliases: []string { "order", "o"},
	Run: func(cmd *cobra.Command, args []string) {
		ListOrders(args)
	},
}
