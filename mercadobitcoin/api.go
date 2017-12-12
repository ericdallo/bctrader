package mercadobitcoin

import (
	"net/http"
	"net/url"
	"encoding/json"
	"time"
	"strings"
	"strconv"
	"log"
	"os"
	"fmt"
)

const (
	NEGOCIATION_API_PATH = "/tapi/v3/"
)

var webClient = &http.Client {
	Timeout: time.Second * 10,
}

func GetOrders() []Order {
	validateCredentials()

	params := url.Values{}
	params.Add("coin_pair", "BRLBTC")
	params.Add("tapi_method", "list_orders")
	params.Add("tapi_nonce", strconv.FormatInt(time.Now().UnixNano(), 10))

	res := requestTapi(params)

	ordersResponse := new(OrdersResponse)

	json.NewDecoder(res.Body).Decode(ordersResponse)

	defer res.Body.Close()

	validateResponse(ordersResponse.StatusCode, ordersResponse.ErrorMessage)

	return ordersResponse.ResponseData.Orders
}

func BuyOrder(quantity string, limitPrice string) Order {
	validateCredentials()

	params := url.Values{}
	params.Add("quantity", quantity)
	params.Add("limit_price", limitPrice)
	params.Add("coin_pair", "BRLBTC")
    params.Add("tapi_method", "place_buy_order")
    params.Add("tapi_nonce", strconv.FormatInt(time.Now().UnixNano(), 10))

    res := requestTapi(params)

    orderResponse := new(OrderResponse)

	json.NewDecoder(res.Body).Decode(orderResponse)

	defer res.Body.Close()

	validateResponse(orderResponse.StatusCode, orderResponse.ErrorMessage)

    return orderResponse.ResponseData.Order
}

func GetPrice(coin string) Price {
	baseurl := BaseURL()

	res, err := webClient.Get(baseurl + "/api/" + coin + "/ticker/")

	if err != nil {
		panic(err)
	}
	price := new(PriceTicker)

	json.NewDecoder(res.Body).Decode(price)

	defer res.Body.Close()

	return price.Ticker
}

func validateCredentials() {
	if !IsCredentialsSetted() {
		log.Printf("tapi_id not configurated, configure with 'bctrader configure [PROVIDER]'")
		os.Exit(-1)
	}
}

func validateResponse(status int, errorMessage string) {
	if errorMessage != "" {
		fmt.Printf("Status: %d \nError: %s", status, errorMessage)
		os.Exit(-1)
	}
}

func requestTapi(params url.Values) *http.Response {
	paramsEncoded := params.Encode()

	negociation_api_url := BaseURL() + NEGOCIATION_API_PATH

	req, err := http.NewRequest("POST", negociation_api_url, strings.NewReader(paramsEncoded))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("TAPI-ID", TapiId())
	req.Header.Set("TAPI-MAC", TapiMacFor(NEGOCIATION_API_PATH + "?" + paramsEncoded))

	resp, err := webClient.Do(req)
	if err != nil {
		panic(err)
	}

	return resp
}