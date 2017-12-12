package cmd

import (
	"github.com/spf13/cobra"
	"github.com/bctrader/mercadobitcoin"
	"fmt"
	"strconv"
	"time"
)

var coin string = "BTC"

func init() {
	rootCmd.AddCommand(priceCmd)
	priceCmd.Flags().StringVarP(&coin, "coin", "c", "", "Coin to show current price")
}

var priceCmd = &cobra.Command {
	Use: 	 "price",
	Short: 	 "Show the current coin price (default bitcoin)",
	Aliases: []string{"$"},
	Run: func(cmd *cobra.Command, args []string) {
		ShowPrice(args)
	},
}

func ShowPrice(args []string) {

	price := mercadobitcoin.GetPrice(coin)

	fmt.Printf("%-15s%-10s\n", "COIN", coin + " (R$)")
	fmt.Printf("%-15s%-10.2f\n", "HIGH", ToNumber(price.High))
	fmt.Printf("%-15s%-10.2f\n", "LOW", ToNumber(price.Low))
	fmt.Printf("%-15s%-10.2f\n", "VOLUME", ToNumber(price.Vol))
	fmt.Printf("%-15s%-10.2f\n", "Last", ToNumber(price.Last))
	fmt.Printf("%-15s%-10.2f\n", "BUY", ToNumber(price.Buy))
	fmt.Printf("%-15s%-10.2f\n", "SELL", ToNumber(price.Sell))
	fmt.Printf("%-15s%-10s\n", "DATE", time.Unix(price.Date, 0))
}

func ToNumber(priceInString string) float64 {
	number, _ := strconv.ParseFloat(priceInString, 64)

	return number;
}