package cmd

import (
	"github.com/spf13/cobra"
	"github.com/bctrader/mercadobitcoin"
)

var purchasesOnly bool
var salesOnly bool

func init() {
	rootCmd.AddCommand(ordersCmd)
	ordersCmd.Flags().BoolVarP(&purchasesOnly, "purchases-only", "b", false, "Show only buyed orders")
	ordersCmd.Flags().BoolVarP(&salesOnly, "sales-only", "s", false, "Show only sold orders")
}

func ListOrders(args[] string) {
	var orderType mercadobitcoin.OrderType

	if purchasesOnly {
		orderType = 1
	} else if salesOnly {
		orderType = 2
	}

	orders := mercadobitcoin.GetOrders(orderType)

	for _, order := range orders {
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