package entity

type OrderQueue struct {
	Orders []*Order
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}

// Less
func (oq *OrderQueue) Less(i int, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

// Swap
func (oq *OrderQueue) Swap(i int, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

// Len
func (oq *OrderQueue) Len() int {
	return len(oq.Orders)
}

// Push
func (oq *OrderQueue) Push(x interface{}) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

// Pop
func (oq *OrderQueue) Pop() interface{} {
	old := oq.Orders
	n := len(old)
	item := old[n-1]

	oq.Orders = old[0 : n-1]

	return item
}
