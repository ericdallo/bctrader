package cmd

import (
	"github.com/spf13/cobra"
	"github.com/bctrader/mercadobitcoin"

	"fmt"
	"os"
)

var sellQuantity string
var sellLimitPrice string

func init() {
	ordersCmd.AddCommand(ordersSellCmd)
	ordersSellCmd.Flags().StringVarP(&sellQuantity, "quantity", "q", "", "Quantity to buy in Ƀ (e.g 0.12)")
	ordersSellCmd.Flags().StringVarP(&sellLimitPrice, "limit-price", "l", "", "Limit price to buy in R$ (e.g 310.00)")
}

func sellOrder(args[] string) {
	order := mercadobitcoin.SellOrder(sellQuantity, sellLimitPrice)

	mercadobitcoin.Print(order)
}

var ordersSellCmd = &cobra.Command {
	Use: 	 "sell",
	Short: 	 "Place a buy order",
	Aliases: []string { "ask", "s"},
	Run: func(cmd *cobra.Command, args []string) {
		if (sellQuantity == "") {
			fmt.Println("You need to specify a 'quantity' in Ƀ (e.g 0.12)")
			os.Exit(-1)
		}

		if (sellLimitPrice == "") {
			fmt.Println("You need to specify a 'limit-price' in R$ (e.g 350.00)")
			os.Exit(-1)
		}

		sellOrder(args)
	},
}
