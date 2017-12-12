package mercadobitcoin

import (
	"fmt"
	"time"
	"strconv"
)

type OrdersResponse struct {
	ResponseData 	OrdersResponseData `json:"response_data"`
	StatusCode      int  			   `json:"status_code"`
	ErrorMessage	string			   `json:"error_message"`
}

type OrdersResponseData struct {
	Orders 				[]Order 	 `json:"orders"`
}

type OrderResponse struct {
	ResponseData OrderResponseData `json:"response_data"`
	StatusCode      int  			   `json:"status_code"`
	ErrorMessage	string			   `json:"error_message"`
}

type OrderResponseData struct {
	Order 				Order 	`json:"order"`
}

type Order struct {
	OrderId          int           `json:"order_id"`
	CoinPair         string        `json:"coin_pair"`
	OrderType        int           `json:"order_type"`
	Status           int           `json:"status"`
	HasFills         bool          `json:"has_fills"`
	Quantity         string        `json:"quantity"`
	LimitPrice       string        `json:"limit_price"`
	ExecutedQuantity string        `json:"executed_quantity"`
	ExecutedPriceAvg string        `json:"executed_price_avg"`
	Fee              string        `json:"fee"`
	CreatedTimestamp int64 	       `json:"created_timestamp"`
	UpdatedTimestamp int64  	   `json:"updated_timestamp"`

	Operations 		 []Operation 	`json:"operations"`
}

type Operation struct {
	OperationId			int		`json:"operation_id"`
	Quantity			string	`json:"quantity"`
	Price				string	`json:"price"`
	FeeRate				string	`json:"fee_rate"`
	ExecutionTimestamp	string	`json:"execution_timestamp"`
}

func Print(order Order) {
	fmt.Printf("%-32s%-10d\n", "Id", order.OrderId)
	fmt.Printf("\t\t%-30s%-10s\n", "Coin Pair", order.CoinPair)
	fmt.Printf("\t\t%-30s%-10s\n", "Type", OrderType(order.OrderType).String())
	fmt.Printf("\t\t%-30s%-10s\n", "Status", OrderStatus(order.Status).String())
	fmt.Printf("\t\t%-30s%-10t\n", "Has fills", order.HasFills)
	fmt.Printf("\t\t%-30s%-10s\n", "Quantity (Ƀ)", order.Quantity)
	fmt.Printf("\t\t%-30s%-10.2f\n", "Limit price (R$)", toNumber(order.LimitPrice))
	fmt.Printf("\t\t%-30s%-10s\n", "Executed quantity (Ƀ)", order.ExecutedQuantity)
	fmt.Printf("\t\t%-30s%-10.2f\n", "Executed price avg (R$)", toNumber(order.ExecutedPriceAvg))
	fmt.Printf("\t\t%-30s%-10s\n", "Fee (Ƀ)", order.Fee)
	fmt.Printf("\t\t%-30s%-10s\n", "Created at", time.Unix(order.CreatedTimestamp, 0))
	fmt.Printf("\t\t%-30s%-10s\n", "Updated at", time.Unix(order.UpdatedTimestamp, 0))

	for _, operation := range order.Operations {
		fmt.Printf("\t\t%-30s%-10d\n", "OperationId", operation.OperationId)
		fmt.Printf("\t\t└─ %-30s%-10s\n", "Quantity (Ƀ)", operation.Quantity)
		fmt.Printf("\t\t└─ %-30s%-10s\n", "Price (R$)", operation.Price)
		fmt.Printf("\t\t└─ %-30s%-10s\n", "FeeRate ", operation.FeeRate)
		fmt.Printf("\t\t└─ %-30s%-10s\n", "ExecutionTimestamp", operation.ExecutionTimestamp)
	}
	fmt.Println()
}

func toNumber(priceInString string) float64 {
	number, _ := strconv.ParseFloat(priceInString, 64)

	return number;
}