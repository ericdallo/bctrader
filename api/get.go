package api

import (
	"net/http"
	"encoding/json"
	"time"
	
)

type Ticker struct {
	High      string
    Low		  string
    Vol	  	  string
    Last	  string
    Buy		  string
    Sell 	  string
    Date	  int
}

type Price struct {
	Ticker Ticker
}

var webClient = &http.Client {
	Timeout: time.Second * 10,
}

func GetPrice() Ticker {
	res, err := webClient.Get("https://www.mercadobitcoin.net/api/BTC/ticker/")	

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	price := new(Price)

	json.NewDecoder(res.Body).Decode(price)

	return price.Ticker
}