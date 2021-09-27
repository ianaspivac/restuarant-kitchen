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
		Name:        "John Johnson",
		CatchPhrase: "No time for talking!",
	}
}
func Cooking(nrCooks int) {
	var wg sync.WaitGroup
	wg.Add(nrCooks)
	var m sync.Mutex
	indexOrder := 0
	for i := 0; i < nrCooks; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				if len(Order_list) > 0 {
					m.Lock()
					if len(Order_list) > indexOrder {
						order := Order_list[indexOrder]
						indexOrder++
						m.Unlock()
						fmt.Printf("Cook %v started preparing order\n", i)
						time.Sleep(time.Duration(order.MaxPreparationTime) * time.Second)
						fmt.Printf("Cook %v prepared order\n", i)
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
					}else {
						m.Unlock()
					}
				}
			}
		}(i)
	}
	wg.Wait()

}
