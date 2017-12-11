package api

import (
	"github.com/spf13/viper"
	"net/http"
	"encoding/json"
	"time"
	. "github.com/bctrader/coin"
)

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

var webClient = &http.Client {
	Timeout: time.Second * 10,
}

func GetPrice(coin Coin) Ticker {
	baseurl := viper.GetString("MERCADO_BITCOIN_BASE_URL")

	res, err := webClient.Get(baseurl + "/" + string(coin) + "/ticker/")

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	price := new(Price)

	json.NewDecoder(res.Body).Decode(price)

	return price.Ticker
}