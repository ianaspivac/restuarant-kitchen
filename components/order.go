package components

type Order struct {
	TableId            int     `json:"table_id"`
	OrderId            int     `json:"order_id"`
	WaiterId           int     `json:"waiter_id"`
	Priority           int     `json:"priority"`
	MenuItemIds        []int   `json:"items"`
	MaxPreparationTime float32 `json:"max_wait"`
	PickUpTime         int64   `json:"pick_up_time"`
	Done               bool    `json:"-"`
}

var Order_list []*Order


func InitOrder(order *Order) *Order {
	return &Order{
		TableId:            order.TableId,
		OrderId:            order.OrderId,
		WaiterId:           order.WaiterId,
		Priority:           order.Priority,
		MenuItemIds:        order.MenuItemIds,
		MaxPreparationTime: order.MaxPreparationTime,
		PickUpTime:         order.PickUpTime,
		Done:               false,
	}
}


