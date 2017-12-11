package api

import (
	MercadobitcoinConfig "github.com/bctrader/config"

	"net/http"
	"net/url"
	"encoding/json"
	"time"
	"fmt"
	"strings"
	"io/ioutil"
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

func GetOrder() {
	validateCredentials()

	uri := MercadobitcoinConfig.BaseURL() + NEGOCIATION_API_PATH

	params := url.Values{}
	params.Add("coin_pair", "BRLBTC")
    params.Add("tapi_method", "list_orders")
    params.Add("tapi_nonce", strconv.FormatInt(time.Now().UnixNano(), 10))

    paramsEncoded := params.Encode()

	req, err := http.NewRequest("POST", uri, strings.NewReader(paramsEncoded))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("TAPI-ID", MercadobitcoinConfig.TapiId())
    req.Header.Set("TAPI-MAC", MercadobitcoinConfig.TapiMacFor(NEGOCIATION_API_PATH + "?" + paramsEncoded))

    resp, err := webClient.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}

type Ticker struct {
	High      string
    Low		  string
    Vol	  	  string
    Last	  string
    Buy		  string
    Sell 	  string
    Date	  int64
}

type Price struct {
	Ticker Ticker
}

func GetPrice(coin string) Ticker {
	baseurl := MercadobitcoinConfig.BaseURL()

	res, err := webClient.Get(baseurl + "/api/" + coin + "/ticker/")

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	price := new(Price)

	json.NewDecoder(res.Body).Decode(price)

	return price.Ticker
}

func validateCredentials() {
	if !MercadobitcoinConfig.IsCredentialsSetted() {
		log.Printf("tapi_id not configurated, configure with 'bctrader configure [PROVIDER]'")
		os.Exit(-1)
	}
}