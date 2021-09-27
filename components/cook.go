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

const NrCooks int = 2

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
	for idx := 0; idx < NrCooks; idx++ {
		Cooks = append(Cooks, &Cook{
			Id:          idx,
			Rank:        3,
			Proficiency: 3,
			Name:        "John Johnson",
			CatchPhrase: "No time for talking!",
		})
	}
}
func CooksManagement() {
	for idx, _ := range Cooks {
		go Cooks[idx].Cooking()
	}
}
func getOrderListItem() *Order {
	order := Order_list[0]
	Order_list = Order_list[1:]

	return order
}
var OrderMutex sync.Mutex
func (c *Cook) Cooking() {
	for {
		OrderMutex.Lock()
		if len(Order_list) > 0 {
			order := getOrderListItem()
			OrderMutex.Unlock()

			fmt.Printf("Cook %d started preparing order\n", c.Id)
			time.Sleep(time.Duration(order.MaxPreparationTime) * time.Second)

			fmt.Printf("Cook %d prepared order\n", c.Id)
			fmt.Printf("%+v\n", order)

			jsonBody, err := json.Marshal(order)
			if err != nil {
				log.Panic(err)
			}
			contentType := "application/json"
			_, err = http.Post("http://dining:8080/distribution", contentType, bytes.NewReader(jsonBody))
			if err != nil {
				return
			}
		} else {
			OrderMutex.Unlock()
		}
	}
}
