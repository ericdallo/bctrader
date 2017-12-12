package cmd

import (
	"github.com/spf13/cobra"
	"github.com/bctrader/mercadobitcoin"
)

func init() {
	rootCmd.AddCommand(ordersCmd)
}

func ListOrders(args[] string) {
	for _, order := range mercadobitcoin.GetOrders() {
		mercadobitcoin.Print(order)
	}
}

var ordersCmd = &cobra.Command {
	Use: 	 "orders",
	Short: 	 "List/buy/sell your orders",
	Aliases: []string { "order", "o"},
	Run: func(cmd *cobra.Command, args []string) {
		ListOrders(args)
	},
}