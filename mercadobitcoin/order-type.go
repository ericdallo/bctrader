package mercadobitcoin

type OrderType int

const (
	BUY 		OrderType = 1 + iota
	SELL
)

var alltypes = [...]string {
	"BUY",
	"SELL",
}

func (otype OrderType) String() string {
	return alltypes[otype - 1]
}

func (otype OrderType) Value() int {
	return int(otype)
}