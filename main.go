package kitchen

import (
	"kitchen/components"
	"kitchen/util"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	const nrCooks int = 2
	var cooks [nrCooks]*components.Cook
	//nrApparatus := nrCooks
	for i := 0; i <= nrCooks; i++ {
		cooks[i] = components.HireCook(util.RandomizeNr(3))
	}

}
