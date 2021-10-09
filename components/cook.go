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

func getOrderListItem(rank int) FoodOrder {
	var food FoodOrder
	if rank == 3 && len(FoodList3.GetFoodList()) > 0 {
		idx := findByCookingApparatus(FoodList3.GetFoodList())
		food = FoodList3.GetFoodList()[idx]
		FoodList3.ReduceFoodList(idx)
	} else if rank >= 2 && len(FoodList2.GetFoodList()) > 0 {
		idx := findByCookingApparatus(FoodList2.GetFoodList())
		food = FoodList2.GetFoodList()[idx]
		FoodList2.ReduceFoodList(idx)
	} else {
		idx := findByCookingApparatus(FoodList1.GetFoodList())
		food = FoodList1.GetFoodList()[idx]
		FoodList1.ReduceFoodList(idx)
	}
	FoodToPrepare--
	return food
}

func findByCookingApparatus(foodList []FoodOrder) int {
	if CookingApparatus["oven"] == 0 && CookingApparatus["stove"] == 0 {
		if len(foodList) == 1 && foodList[0].Food.cookingApparatus != ""{
			switch foodList[0].Food.cookingApparatus {
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
			return 0

		}
		for idx, _ := range foodList {
			if foodList[idx].Food.cookingApparatus == "" {
				return idx
			}
		}
	} else {
		for idx, _ := range foodList {
			if foodList[idx].Food.cookingApparatus == "" {
				return idx
			} else if CookingApparatus[foodList[idx].Food.cookingApparatus] != 0 {
				switch foodList[idx].Food.cookingApparatus {
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
				return idx
			}
		}
	}

	return 0
}

func addToFinishedFoods(food FoodOrder) {
	if _, alreadyExists := ReadyFood[food.orderId]; alreadyExists {
		ReadyFood[food.orderId]++
	} else {
		ReadyFood[food.orderId] = 1
	}
	if ReadyFood[food.orderId] == food.orderSize {
		for idx, _ := range Order_list {
			if Order_list[idx].OrderId == food.orderId {
				fmt.Printf("Order was prepared: %+v\n", Order_list[idx])
				Order_list = Order_list[1:]
				jsonBody, err := json.Marshal(Order_list[idx])
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

			OrderMutex.Unlock()

			fmt.Printf("-Cook %d started preparing food %+v\n", c.Id, food)
			<-time.After(time.Duration(food.preparationTime) * time.Second)

			CookingApparatus[food.Food.cookingApparatus] += 1
			if food.Food.cookingApparatus == "oven" {
				<-semOven
			} else if food.Food.cookingApparatus == "stove" {
				<-semStove
			}

			fmt.Printf("+Cook %d finished preparing food %+v\n", c.Id, food)
			addToFinishedFoods(food)

		} else {
			OrderMutex.Unlock()
		}
	}
}
