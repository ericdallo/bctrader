package cmd

import (
	"github.com/spf13/cobra"
	"github.com/bctrader/mercadobitcoin"

	"fmt"
	"os"
)

var buyQuantity string
var buyLimitPrice string

func init() {
	ordersCmd.AddCommand(ordersBuyCmd)
	ordersBuyCmd.Flags().StringVarP(&buyQuantity, "quantity", "q", "", "Quantity to buy in Ƀ (e.g 0.12)")
	ordersBuyCmd.Flags().StringVarP(&buyLimitPrice, "limit-price", "l", "", "Limit price to buy in R$ (e.g 310.00)")
}

func BuyOrder(args[] string) {
	order := mercadobitcoin.BuyOrder(buyQuantity, buyLimitPrice)

	mercadobitcoin.Print(order)
}

var ordersBuyCmd = &cobra.Command {
	Use: 	 "buy",
	Short: 	 "Place a buy order",
	Aliases: []string { "purchase", "bid", "b"},
	Run: func(cmd *cobra.Command, args []string) {
		if buyQuantity == "" || buyLimitPrice == "" {
			fmt.Println("You need to specify a 'quantity' (e.g 0.12) and a 'limit-price' (e.g 350.00)")
			os.Exit(-1)
		}

		BuyOrder(args)
	},
}
