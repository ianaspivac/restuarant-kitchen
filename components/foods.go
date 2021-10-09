package components

import (
	"fmt"
	"sync"
)

type Food struct {
	id               int
	name             string
	preparationTime  float32
	complexity       int
	cookingApparatus string
}
type FoodLists struct {
	foodListMutex sync.Mutex
	foodList []FoodOrder

}
var FoodList1 FoodLists
var FoodList2 FoodLists
var FoodList3 FoodLists
var FoodToPrepare = 0
var ReadyFood = make(map[int]int)


type FoodOrder struct {
	orderId       int
	orderSize     int
	orderPriority int
	Food
}


func (f *FoodLists) GetFoodList() []FoodOrder {
	f.foodListMutex.Lock()
	defer f.foodListMutex.Unlock()

	return f.foodList
}

func (f *FoodLists) ReduceFoodList() {
	f.foodListMutex.Lock()
	defer f.foodListMutex.Unlock()

	f.foodList = f.foodList[1:]
}

func (f *FoodLists) SetFoodList(order *Order,idx int) {
	f.foodListMutex.Lock()
	defer f.foodListMutex.Unlock()

	f.foodList = append(f.foodList,FoodOrder{
		orderId:       order.OrderId,
		orderSize:     len(order.MenuItemIds),
		orderPriority: order.Priority,
		Food:          Menu[idx],
	})
}

func SeparateFoods(order *Order) {
	for _, val := range order.MenuItemIds {
		switch Menu[val-1].complexity {
		case 1:
			FoodList1.SetFoodList(order,val-1)
		case 2:
			FoodList2.SetFoodList(order,val-1)
		case 3:
			FoodList3.SetFoodList(order,val-1)
		default:
			fmt.Printf("Unexpected complexity")
		}
		FoodToPrepare++
	}

}

var Menu = []Food{
	{
		id:               1,
		name:             "pizza",
		preparationTime:  20,
		complexity:       2,
		cookingApparatus: "oven",
	},
	{
		id:               2,
		name:             "salad",
		preparationTime:  10,
		complexity:       1,
		cookingApparatus: "",
	},
	{
		id:               3,
		name:             "zeama",
		preparationTime:  7,
		complexity:       1,
		cookingApparatus: "stove",
	},
	{
		id:               4,
		name:             "Scallop Sashimi with Meyer Lemon Confit",
		preparationTime:  32,
		complexity:       3,
		cookingApparatus: "",
	},
	{
		id:               5,
		name:             "Island Duck with Mulberry Mustard",
		preparationTime:  35,
		complexity:       3,
		cookingApparatus: "oven",
	},
	{
		id:               6,
		name:             "Waffles",
		preparationTime:  10,
		complexity:       1,
		cookingApparatus: "stove",
	},
	{
		id:               7,
		name:             "Aubergine",
		preparationTime:  20,
		complexity:       2,
		cookingApparatus: "",
	},
	{
		id:               8,
		name:             "Lasagna",
		preparationTime:  30,
		complexity:       2,
		cookingApparatus: "oven",
	},
	{
		id:               9,
		name:             "Burger",
		preparationTime:  15,
		complexity:       1,
		cookingApparatus: "oven",
	},
	{
		id:               10,
		name:             "Gyros",
		preparationTime:  15,
		complexity:       1,
		cookingApparatus: "",
	},
}
