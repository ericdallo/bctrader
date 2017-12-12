package mercadobitcoin

type OrderResponse struct {
	ResponseData ResponseData `json:"response_data"`
}

type ResponseData struct {
	Orders 		[]Order 	 `json:"orders"`
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
