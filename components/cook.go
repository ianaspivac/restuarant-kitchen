package components

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

const NrCooks int = 3

var Cooks []*Cook
var CookingApparatus map[string]int
var semStove = make(chan int, 2)
var semOven = make(chan int, 1)

type Cook struct {
	Id          int
	Rank        int
	Proficiency int
	Name        string
	CatchPhrase string
}
type Rank int

func HireCooks() {
	for idx := 1; idx <= NrCooks; idx++ {
		Cooks = append(Cooks, &Cook{
			Id:          idx,
			Rank:        idx,
			Proficiency: idx,
			Name:        "John Johnson",
			CatchPhrase: "No time for talking!",
		})
	}
}
func CooksManagement() {
	for idx, _ := range Cooks {
		for foodsAtTime := 0; foodsAtTime < Cooks[idx].Proficiency; foodsAtTime++ {
			go Cooks[idx].Cooking()
		}
	}
}

//logic for cooks rank and proficiency to prepare food
func getOrderListItem(rank int) FoodOrder {
	var food FoodOrder
	if rank == 3 && len(FoodList3.GetFoodList()) > 0 {
		food = findByCookingApparatus(FoodList3.GetFoodList(), rank)
	} else if rank >= 2 && len(FoodList2.GetFoodList()) > 0 {
		food = findByCookingApparatus(FoodList2.GetFoodList(), rank)
	} else {
		food = findByCookingApparatus(FoodList1.GetFoodList(), rank)
	}
	return food
}

//logic for cooking apparatus
func findByCookingApparatus(foodList []FoodOrder, rank int) FoodOrder {
	var food FoodOrder

	if len(foodList) == 1 && CookingApparatus["oven"] == 0 && foodList[0].Food.cookingApparatus == "oven" {
		food = selectListByRank(rank, 0)
		semOven <- 1
		CookingApparatus["oven"] -= 1
		return food
	} else if len(foodList) == 1 && CookingApparatus["stove"] == 0 && foodList[0].Food.cookingApparatus == "stove" {
		food = selectListByRank(rank, 0)
		semStove <- 1
		CookingApparatus["stove"] -= 1
		return food
	} else {
		for idx := range foodList {
			if foodList[idx].Food.cookingApparatus == "" {
				return selectListByRank(rank, idx)
			} else if CookingApparatus[foodList[idx].Food.cookingApparatus] != 0 {
				food = selectListByRank(rank, idx)
				switch food.Food.cookingApparatus {
				case "oven":
					{
						semOven <- 1
						CookingApparatus["oven"] -= 1
					}
				case "stove":
					{
						semStove <- 1
						CookingApparatus["stove"] -= 1
					}
				}
				return food
			}else{
				food = selectListByRank(rank, 0)
				switch food.Food.cookingApparatus {
				case "oven":
					{
						semOven <- 1
						CookingApparatus["oven"] -= 1
					}
				case "stove":
					{
						semStove <- 1
						CookingApparatus["stove"] -= 1
					}
				}
				return food
			}
		}
	}
	return food
}

func selectListByRank(rank int, idx int) FoodOrder {
	var food FoodOrder
	if rank == 3 && len(FoodList3.GetFoodList()) > 0 {
		food = FoodList3.GetFoodList()[idx]
		FoodList3.ReduceFoodList(idx)
	} else if rank >= 2 && len(FoodList2.GetFoodList()) > 0 {
		food = FoodList2.GetFoodList()[idx]
		FoodList2.ReduceFoodList(idx)
	} else {
		food = FoodList1.GetFoodList()[idx]
		FoodList1.ReduceFoodList(idx)
	}
	FoodToPrepare--
	defer OrderMutex.Unlock()
	return food
}

//sending prepared order
func addToFinishedFoods(food FoodOrder, cookId int) {
	for idx := range ReadyFoodsList {
		if food.orderId == ReadyFoodsList[idx].GetOrderIdReadyFoods() {
			ReadyFoodsList[idx].AppendPreparedFood(food.id, cookId)
			if ReadyFoodsList[idx].GetOrderSizeReadyFoods() == food.orderSize {
				for idx, _ := range Order_list {
					if Order_list[idx].OrderId == food.orderId {
						orderPrepared := &OrderPrepared{
							Order:          *Order_list[idx],
							CookingTime:    time.Now().Unix() - Order_list[idx].PickUpTime,
							CookingDetails: ReadyFoodsList[idx].GetListReadyFoods(),
						}
						fmt.Printf("Prepared order: %+v\n", orderPrepared)
						Order_list = Order_list[1:]
						jsonBody, err := json.Marshal(orderPrepared)
						if err != nil {
							log.Panic(err)
						}
						contentType := "application/json"
						_, err = http.Post("http://localhost:8080/distribution", contentType, bytes.NewReader(jsonBody))
						if err != nil {
							return
						}
						break
					}
				}
			}
			break
		}
	}
}

var OrderMutex sync.Mutex

func (c *Cook) Cooking() {
	CookingApparatus = map[string]int{
		"oven":  1,
		"stove": 2,
	}
	for {
		OrderMutex.Lock()
		if (c.Rank == 1 && len(FoodList1.GetFoodList()) < 1) || (c.Rank == 2 && len(FoodList2.GetFoodList()) < 1) {
			OrderMutex.Unlock()
			continue
		}
		if FoodToPrepare > 0 {

			food := getOrderListItem(c.Rank)
			fmt.Printf("-Cook %d started preparing food %+v\n", c.Id, food)

			<-time.After(time.Duration(food.preparationTime) * time.Second)

			fmt.Printf("+Cook %d finished preparing food %+v\n", c.Id, food)

			CookingApparatus[food.Food.cookingApparatus] += 1
			if food.Food.cookingApparatus == "oven" {
				<-semOven
			} else if food.Food.cookingApparatus == "stove" {
				<-semStove
			}

			addToFinishedFoods(food, c.Id)

		} else {
			OrderMutex.Unlock()
		}
	}
}
