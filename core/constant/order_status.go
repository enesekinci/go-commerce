package constant

type OrderStatus uint8

const (
	Canceled       OrderStatus = 1
	Completed      OrderStatus = 2
	OnHold         OrderStatus = 3
	Pending        OrderStatus = 4
	PendingPayment OrderStatus = 5

	Processing OrderStatus = 6
	Refunded   OrderStatus = 7
)
