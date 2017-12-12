package mercadobitcoin

type Price struct {
	High		string
	Low			string
	Vol			string
	Last		string
	Buy			string
	Sell		string
	Date		int64
}

type PriceTicker struct {
	Ticker Price
}