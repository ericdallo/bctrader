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
)

const (
	NEGOCIATION_API_PATH = "/tapi/v3/"
)

var webClient = &http.Client {
	Timeout: time.Second * 10,
}

func GetOrders() []Order {
	validateCredentials()

	uri := BaseURL() + NEGOCIATION_API_PATH

	params := url.Values{}
	params.Add("coin_pair", "BRLBTC")
    params.Add("tapi_method", "list_orders")
    params.Add("tapi_nonce", strconv.FormatInt(time.Now().UnixNano(), 10))

    paramsEncoded := params.Encode()

	req, err := http.NewRequest("POST", uri, strings.NewReader(paramsEncoded))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("TAPI-ID", TapiId())
    req.Header.Set("TAPI-MAC", TapiMacFor(NEGOCIATION_API_PATH + "?" + paramsEncoded))

    resp, err := webClient.Do(req)
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    orderResponse := new(OrderResponse)

	json.NewDecoder(resp.Body).Decode(orderResponse)

    return orderResponse.ResponseData.Orders
}

func GetPrice(coin string) Price {
	baseurl := BaseURL()

	res, err := webClient.Get(baseurl + "/api/" + coin + "/ticker/")

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	price := new(PriceTicker)

	json.NewDecoder(res.Body).Decode(price)

	return price.Ticker
}

func validateCredentials() {
	if !IsCredentialsSetted() {
		log.Printf("tapi_id not configurated, configure with 'bctrader configure [PROVIDER]'")
		os.Exit(-1)
	}
}