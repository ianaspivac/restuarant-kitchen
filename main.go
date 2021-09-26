package main

import (
	"github.com/gin-gonic/gin"
	"kitchen/components"
	"kitchen/util"
	"math/rand"
	"net/http"
	"time"
)



func getOrder(c *gin.Context) {
	var order *components.Order
	if err := c.BindJSON(&order); err != nil {
		return
	}
	components.Order_list = append(components.Order_list, components.InitOrder(order))
	c.IndentedJSON(http.StatusCreated, order)
}

func main() {

	router := gin.Default()
	router.POST("/order", getOrder)

	rand.Seed(time.Now().UnixNano())

	const nrCooks int = 2
	var cooks [nrCooks]*components.Cook
	//nrApparatus := nrCooks
	for i := 0; i < nrCooks; i++ {
		cooks[i] = components.HireCook(util.RandomizeNr(3))
	}
	go func() { components.Cooking(nrCooks) }()
	router.Run("localhost:8081")
}
