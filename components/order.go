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

func UpdateOrder(order *Order) *Order {
	return &Order{
		TableId:            order.TableId,
		OrderId:            order.OrderId,
		WaiterId:           order.WaiterId,
		Priority:           order.Priority,
		MenuItemIds:        order.MenuItemIds,
		MaxPreparationTime: order.MaxPreparationTime,
		PickUpTime:         order.PickUpTime,
		Done:               true,
	}
}
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

func GetOrderPrepTime() (float32, int) {
	if len(Order_list) > 0 {
		for i := 0; i <= len(Order_list); i++ {
			if Order_list[i].Done == false {
				return Order_list[i].MaxPreparationTime, 0
			}
		}
	}
	return 0, 0
}
