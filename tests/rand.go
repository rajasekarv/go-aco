package main

import (
	"experiments/ant_colony_optimization/weightedrandom"
	"fmt"
)

func main() {
	choices := make([]weightedrandom.Choice, 0, 20)

	//fmt.Println(choices)
	choices = append(choices, weightedrandom.Choice{0.34, 0.34})
	choices = append(choices, weightedrandom.Choice{0.6, 0.6})
	choices = append(choices, weightedrandom.Choice{0.06, 0.06})
	//fmt.Println(choices) // [{1 dg} {2 n}]
	count1 := 0
	count2 := 0
	count3 := 0
	for i := 0; i < 100000; i++ {
		result := weightedrandom.WeightedChoice(choices)
		//if err != nil {
		//	panic(err)
		//}

		//fmt.Println("here ---> ", result.Item) //{2 n}
		if result.Weight == 0.34 {
			count1++
		} else if result.Weight == 0.6 {
			count2++
		} else {
			count3++
		}
		//fmt.Println("choice1 ", count1, "choice2 ", count1)
	}
	fmt.Println("choice1 ", count1, "choice2 ", count2, "choice3", count3)
}
