package cmd

import (
	"github.com/spf13/cobra"
	"github.com/bctrader/api"
	Coin "github.com/bctrader/coin"
	"fmt"
	"time"
)

func init() {
	rootCmd.AddCommand(priceCmd)
}

var priceCmd = &cobra.Command {
	Use: 	 "price",
	Short: 	 "Show the current coin price (default bitcoin)",
	Aliases: []string{"$"},
	Run: func(cmd *cobra.Command, args []string) {
		showPrice(args)
	},
}

func showPrice(args []string) {

	coin := Coin.BTC
	
	price := api.GetPrice(coin)

	fmt.Printf("%-15s%-10s\n", "COIN", coin + " (R$)")
	fmt.Printf("%-15s%-10s\n", "HIGH", price.High)
	fmt.Printf("%-15s%-10s\n", "LOW", price.Low)
	fmt.Printf("%-15s%-10s\n", "VOLUME", price.Vol)
	fmt.Printf("%-15s%-10s\n", "Last", price.Last)
	fmt.Printf("%-15s%-10s\n", "BUY", price.Buy)
	fmt.Printf("%-15s%-10s\n", "SELL", price.Sell)
	fmt.Printf("%-15s%-10s\n", "DATE", time.Unix(price.Date, 0))
}