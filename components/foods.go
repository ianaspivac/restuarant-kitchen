package components

import (
	"fmt"
	"sort"
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
	foodList      []FoodOrder
}

type CookFood struct {
	FoodId int `json:"food_id"`
	CookId int `json:"cook_id"`
}

type ReadyFoods struct {
	orderId   int
	orderSize int
	foods     []CookFood
}

func (f *ReadyFoods) AppendPreparedFood(foodId int, cookId int) {
	f.orderSize++
	f.foods = append(f.foods, CookFood{
		FoodId: foodId,
		CookId: cookId,
	})
}
func InitReadyFoods(order *Order) {
	ReadyFoodsList = append(append(ReadyFoodsList, ReadyFoods{
		orderId: order.OrderId,
		orderSize:0,
	}))

}
func (f *ReadyFoods) GetOrderIdReadyFoods() int {
	return f.orderId
}
func (f *ReadyFoods) GetOrderSizeReadyFoods() int {
	return f.orderSize
}
func (f *ReadyFoods) GetListReadyFoods() []CookFood {
	return f.foods
}
var FoodList1 FoodLists
var FoodList2 FoodLists
var FoodList3 FoodLists
var FoodToPrepare = 0

var ReadyFoodsList []ReadyFoods

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

func (f *FoodLists) SetFoodList(order *Order, idx int) {
	f.foodListMutex.Lock()
	defer f.foodListMutex.Unlock()

	f.foodList = append(f.foodList, FoodOrder{
		orderId:       order.OrderId,
		orderSize:     len(order.MenuItemIds),
		orderPriority: order.Priority,
		Food:          Menu[idx],
	})
	f.SortFoodList()
	f.AddPriority(order.Priority)
}

func RemoveIndex(s []FoodOrder, index int) []FoodOrder {
	return append(s[:index], s[index+1:]...)
}

func (f *FoodLists) ReduceFoodList(idx int) {
	f.foodListMutex.Lock()
	defer f.foodListMutex.Unlock()

	f.foodList = RemoveIndex(f.foodList, idx)
}

//sorting according to priority whitout changing items if priority is equal
func (f *FoodLists) SortFoodList() {
	sort.SliceStable(f.foodList, func(i, j int) bool {
		return f.foodList[i].orderPriority > f.foodList[j].orderPriority
	})
}

func findMin(foodList []FoodOrder) int {
	min := 5
	for _, val := range foodList {
		if val.orderPriority < min {
			min = val.orderPriority
		}

	}
	return min
}

//to not get in situation where the lowest priority is never cooked
// every time diference between added order priority and min priority is > 2
// all items with priority <5 will have incremented priority by 1

func (f *FoodLists) AddPriority(addedOrderPriority int) {
	if (addedOrderPriority - findMin(f.foodList)) > 2 {
		for _, val := range f.foodList {
			if val.orderPriority != 5 {
				val.orderPriority += 1
			}
		}
	}
}

func SeparateFoods(order *Order) {
	for _, val := range order.MenuItemIds {
		switch Menu[val-1].complexity {
		case 1:
			FoodList1.SetFoodList(order, val-1)
		case 2:
			FoodList2.SetFoodList(order, val-1)
		case 3:
			FoodList3.SetFoodList(order, val-1)
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
