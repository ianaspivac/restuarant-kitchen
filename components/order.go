package components

type Order struct {
	TableId            int     `json:"table_id"`
	OrderId            int     `json:"order_id"`
	Priority           int     `json:"priority"`
	MenuItemIds        []int   `json:"items"`
	MaxPreparationTime float32 `json:"max_wait"`
	PickUpTime         int64   `json:"pick_up_time"`
}