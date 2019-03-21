package main

import (
	"bufio"
	"experiments/ant_colony_optimization/acoseq"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

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
		//fmt.Println(line)
		for _, str := range line[1:] {
			//fmt.Println(str)
			dis, _ := strconv.ParseFloat(strings.Trim(str, " "), 10)
			//fmt.Println(dis)
			//fmt.Println(line[i])
			if dis == 0.0 {
				row = append(row, math.Inf(1))
			} else {
				row = append(row, dis)

			}
		}
		//fmt.Println(line, len(line))
		//fmt.Println(row, len(row))
		distances = append(distances, row)
		//break
	}
	//fmt.Println(distances, len(distances), len(distances[0]))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var antColony acoseq.AntColony
	//antColony = acoseq.AntColony{200, 190, 200, 0.85, 1, 1, distances}
	//fmt.Println(distances)
	//fmt.Println(distances[0:25][0:25])
	antColony = acoseq.AntColony{
		NoAnts:      1,
		NoBestPath:  1,
		NoIteration: 1,
		Decay:       0.85,
		Alpha:       1,
		Beta:        1,
		Distances:   distances,
	}

	shortestPath := antColony.Run()
	fmt.Println(shortestPath)

}
