package cmd

import (
	"github.com/spf13/cobra"
	"github.com/bctrader/mercadobitcoin"

	"fmt"
	"time"
)

func init() {
	rootCmd.AddCommand(ordersCmd)
}

func ListOrders(args[] string) {
	fmt.Println("Orders list")
	for _, order := range mercadobitcoin.GetOrders() {
		fmt.Printf("---%-20s%-10d\n", "Id", order.OrderId)
		fmt.Printf("---%-20s%-10s\n", "Coin Pair", order.CoinPair)
		fmt.Printf("---%-20s%-10d\n", "Type", order.OrderType)
		fmt.Printf("---%-20s%-10d\n", "Status", order.Status)
		fmt.Printf("---%-20s%-10b\n", "Has fills", order.HasFills)
		fmt.Printf("---%-20s%-10s\n", "Quantity", order.Quantity)
		fmt.Printf("---%-20s%-10s\n", "Limit price", order.LimitPrice)
		fmt.Printf("---%-20s%-10s\n", "Executed quantity", order.ExecutedQuantity)
		fmt.Printf("---%-20s%-10s\n", "Executed price avg", order.ExecutedPriceAvg)
		fmt.Printf("---%-20s%-10s\n", "Fee", order.Fee)
		fmt.Printf("---%-20s%-10s\n", "Created at", time.Unix(order.CreatedTimestamp, 0))
		fmt.Printf("---%-20s%-10s\n", "Updated at", time.Unix(order.UpdatedTimestamp, 0))

		for _, operation := range order.Operations {
			fmt.Printf("------%-20s%-10s\n", "OperationId", operation.OperationId)
			fmt.Printf("------%-20s%-10s\n", "Quantity", operation.Quantity)
			fmt.Printf("------%-20s%-10s\n", "Price", operation.Price)
			fmt.Printf("------%-20s%-10s\n", "FeeRate", operation.FeeRate)
			fmt.Printf("------%-20s%-10s\n", "ExecutionTimestamp", operation.ExecutionTimestamp)
			fmt.Println()
		}
	}
}

var ordersCmd = &cobra.Command {
	Use: 	 "orders",
	Short: 	 "List all your orders",
	Aliases: []string { "order", "o"},
	Run: func(cmd *cobra.Command, args []string) {
		ListOrders(args)
	},
}
