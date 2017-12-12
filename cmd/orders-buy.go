package cmd

import (
	"github.com/spf13/cobra"
	"github.com/bctrader/mercadobitcoin"

	"fmt"
	"os"
)

var quantity string
var limitPrice string

func init() {
	ordersCmd.AddCommand(ordersBuyCmd)
	ordersBuyCmd.Flags().StringVarP(&quantity, "quantity", "q", "", "Quantity to buy (e.g 300.00)")
	ordersBuyCmd.Flags().StringVarP(&limitPrice, "limit-price", "l", "", "Limit price to buy (e.g 310.00)")
}

func BuyOrder(args[] string) {
	order := mercadobitcoin.BuyOrder(quantity, limitPrice)

	mercadobitcoin.Print(order)
}

var ordersBuyCmd = &cobra.Command {
	Use: 	 "buy",
	Short: 	 "Place a buy order",
	Aliases: []string { "purchase", "bid", "b"},
	Run: func(cmd *cobra.Command, args []string) {
		if (quantity == "") {
			fmt.Println("You need to specify a 'quantity' (e.g 330.10)")
			os.Exit(-1)
		}

		if (limitPrice == "") {
			fmt.Println("You need to specify a 'limit-price' (e.g 350.00)")
			os.Exit(-1)
		}

		BuyOrder(args)
	},
}
