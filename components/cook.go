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
		food = FoodList3.GetFoodList()[0]
		FoodList3.ReduceFoodList()
	} else if rank >= 2 && len(FoodList2.GetFoodList()) > 0 {
		food = FoodList2.GetFoodList()[0]
		FoodList2.ReduceFoodList()
	} else {
		food = FoodList1.GetFoodList()[0]
		FoodList1.ReduceFoodList()
	}
	FoodToPrepare--
	return food
}
func addToFinishedFoods(food FoodOrder) {
	if _, alreadyExists := ReadyFood[food.orderId]; alreadyExists {
		ReadyFood[food.orderId]++
	} else {
		ReadyFood[food.orderId] = 1
	}
	if ReadyFood[food.orderId] == food.orderSize {
		//feel like need some semaphore or mutex
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
	//CookingAparatus := map[string]int{
	//	"oven":2,
	//	"stove":1,
	//}
	for {
		OrderMutex.Lock()
		if (c.Rank == 1 && len(FoodList1.GetFoodList()) < 1)  || (c.Rank == 2 && len(FoodList2.GetFoodList()) < 1){
			OrderMutex.Unlock()
			continue
		}
		if FoodToPrepare > 0 {
			//if(CookingAparatus["oven"] == 0 && CookingAparatus["stove"] == 0)
			food := getOrderListItem(c.Rank)

			OrderMutex.Unlock()

			fmt.Printf("-Cook %d started preparing food %+v\n", c.Id, food)
			<-time.After(time.Duration(food.preparationTime) * time.Second)
			fmt.Printf("+Cook %d finished preparing food %+v\n", c.Id, food)
			addToFinishedFoods(food)

		} else {
			OrderMutex.Unlock()
		}
	}
}
