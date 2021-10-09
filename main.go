package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kitchen/components"
	"math/rand"
	"net/http"
	"time"
)



func getOrder(c *gin.Context) {
	var order *components.Order
	if err := c.BindJSON(&order); err != nil {
		return
	}

	fmt.Printf("Recieved order to cook: %+v \n",order)
	components.SeparateFoods(order)
	components.Order_list = append(components.Order_list, order)
	c.IndentedJSON(http.StatusCreated, order)
}

func main() {
	router := gin.Default()
	router.POST("/order", getOrder)

	rand.Seed(time.Now().UnixNano())

	components.HireCooks()
	components.CooksManagement()
	router.Run(":8081")
}
