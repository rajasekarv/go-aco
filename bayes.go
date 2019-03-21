package main

import (
	"log"
	"math"

	"github.com/d4l3k/go-bayesopt"
)

func main() {
	X := bayesopt.UniformParam{
		Max: 10,
		Min: -10,
	}
	o := bayesopt.New(
		[]bayesopt.Param{
			X,
		},
	)
	// minimize x^2+1
	x, y, err := o.Optimize(func(params map[bayesopt.Param]float64) float64 {
		return math.Pow(params[X], 2) + 1
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(x, y)
}
