package util

import (
	"math/rand"
)

func RandomizeNr(max int)int{
	return (rand.Intn(max) + 1)
}