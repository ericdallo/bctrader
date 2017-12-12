package mercadobitcoin

type OrderStatus int

const (
	OPEN 		OrderStatus = 2
	CANCELED	OrderStatus = 3
	FILLED		OrderStatus = 4
)

var allstatus = [...]string {
	"OPEN",
	"CANCELED",
	"FILLED",
}

func (status OrderStatus) String() string {
	return allstatus[status - 2]
}