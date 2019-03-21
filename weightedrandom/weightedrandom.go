package weightedrandom

import (
	"math/rand"
	"time"
)

func randFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	r := min + rand.Float64()*(max-min)
	//fmt.Println(r)
	return r
}

type Choice struct {
	Weight float64
	Item   interface{}
}

func WeightedChoice(choices []Choice) Choice {
	var ret Choice
	sum := 0.0
	for _, c := range choices {
		sum += c.Weight
	}
	r := randFloat(0, sum)
	for _, c := range choices {
		r -= c.Weight
		if r < 0 {
			return c
		}
	}
	return ret
}
