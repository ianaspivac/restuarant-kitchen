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

type Cook struct {
	Rank        int
	Proficiency int
	Name        string
	CatchPhrase string
}
type Rank int

func HireCook(rank int) *Cook {
	return &Cook{
		Rank:        rank,
		Proficiency: rank,
		Name: "John Johnson",
		CatchPhrase: "No time for talking!",
	}
}
func Cooking(nrCooks int){
	var wg sync.WaitGroup
	wg.Add(nrCooks)
	var m sync.Mutex
	for i := 0; i < nrCooks; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				m.Lock()
				orderPrepTime,orderIndex := GetOrderPrepTime()
				if(orderPrepTime != 0) {
					time.Sleep(time.Duration(orderPrepTime) * time.Second)
					order := UpdateOrder(Order_list[orderIndex])
					fmt.Printf("Cook %v prepared order\n",i)
					fmt.Printf("%+v\n", order)
					jsonBody, err := json.Marshal(order)
					if err != nil {
						log.Panic(err)
					}
					contentType := "application/json"
					_, err = http.Post("http://localhost:8080/distribution", contentType, bytes.NewReader(jsonBody))
					if err != nil {
						return
					}
				}
				m.Unlock()

			}
		}(i)
	}
	wg.Wait()

}
