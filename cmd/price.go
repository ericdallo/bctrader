package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/bctrader/api"
)

func showPrice(args []string) {
	
	price := api.GetPrice()

	fmt.Printf("%-10s%-15s%-15s%-20s%-15s%-15s%-20s%-10s\n", "COIN", "HIGH", "LOW", "VOLUME", "LAST", "BUY", "SELL", "DATE")

	fmt.Printf("%-10s%-15s%-15s%-20s%-15s%-15s%-20s%-10d\n", "BTC", price.High, price.Low, price.Vol, price.Last, price.Buy, price.Sell, price.Date)
}

var priceCmd = &cobra.Command {
	Use: 	 "price",
	Short: 	 "Show the current bitcoin price",
	Aliases: []string{"$"},
	Run: func(cmd *cobra.Command, args []string) {
		showPrice(args)
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)
}