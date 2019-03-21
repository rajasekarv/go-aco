package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func shortestTour(nodes [][]float64, presentNode int, visitedNodes int, lookupMatrix [][]float64) float64 {
	if visitedNodes == (1<<uint(len(nodes)) - 1) {
		return nodes[presentNode][0]
	}

	if lookupMatrix[presentNode][visitedNodes] != math.Inf(1) {
		return lookupMatrix[presentNode][visitedNodes]
	}
	//var tour []int
	for i := 0; i < len(nodes); i++ {
		if (i == presentNode) || ((visitedNodes & (1 << uint(i))) > 0) {
			continue
		}

		newDistance := nodes[presentNode][i] + shortestTour(nodes, i, visitedNodes|(1<<uint(i)), lookupMatrix)

		if newDistance < lookupMatrix[presentNode][visitedNodes] {
			//tour = append(tour, i)
			lookupMatrix[presentNode][visitedNodes] = newDistance
		}
	}
	//fmt.Println("shortestTour ", tour)
	return lookupMatrix[presentNode][visitedNodes]
}

func main() {
	var distances [][]float64
	file, err := os.Open("/home/raja/Downloads/programming/python/AntColonyOptimization-master/att48_d.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row []float64
		line := strings.Split(scanner.Text(), "      ")
		for _, str := range line[1:] {
			dis, _ := strconv.ParseFloat(strings.Trim(str, " "), 10)
			if dis == 0.0 {
				row = append(row, math.Inf(1))
			} else {
				row = append(row, dis)

			}
		}
		distances = append(distances, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//distances1 := [][]float64{{0, 20, 42, 25},
	//	{20, 0, 30, 34},
	//	{42, 30, 0, 10},
	//	{25, 34, 10, 0}}
	var distances1 [][]float64
	var lookupMatrix [][]float64
	for i := 0; i < 25; i++ {
		var temp []float64
		for j := 0; j < 25; j++ {
			temp = append(temp, distances[i][j])
		}
		distances1 = append(distances1, temp)
	}
	for i := 0; i < len(distances1); i++ {
		var temp []float64
		for j := 0; j < (1<<uint(len(distances1)))-1; j++ {
			temp = append(temp, math.Inf(1))
		}
		lookupMatrix = append(lookupMatrix, temp)
	}
	fmt.Println("nodes size:", len(distances1), " ", len(distances1[0]))
	fmt.Println("lookup size:", len(lookupMatrix), " ", len(lookupMatrix[0]))
	fmt.Println("shortest distance", shortestTour(distances1, 0, 1, lookupMatrix))
}
