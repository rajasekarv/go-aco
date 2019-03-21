package main

import (
	"bufio"
	"experiments/ant_colony_optimization/weightedrandom"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/set"
)

//Ant colony struct
type AntColony struct {
	noAnts      int
	noBestPath  int
	noIteration int
	decay       float64
	alpha       float64
	beta        float64
	distances   [][]float64
	nodes       []int
}

type unitPath struct {
	start int
	end   int
}

var (
	pheromones [][]float64
)

type TourPath struct {
	path     []unitPath
	distance float64
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = float64(int(MaxUint >> 1))

//Run() performs the ant colony optimization and returns the shortest tour path among all iterations.
func (a *AntColony) Run() TourPath {
	//fmt.Println(MaxInt)
	for i := 0; i < len(a.distances); i++ {
		a.nodes = append(a.nodes, i)
	}

	for row := 0; row < len(a.distances); row++ {
		var rowValues []float64
		for value := 0; value < len(a.distances[row]); value++ {
			rowValues = append(rowValues, 1.0/float64(len(a.distances)))
		}
		pheromones = append(pheromones, rowValues)
	}

	var shortestPathIteration TourPath
	var shortestPathAllTime TourPath
	shortestPathAllTime.distance = MaxInt
	for iteration := 0; iteration < a.noIteration; iteration++ {
		var iterationPaths []TourPath
		iterationPaths = a.genPaths()
		a.spreadPheromone(iterationPaths, a.noBestPath)
		shortestPathIteration = iterationPaths[0]
		for _, iterationPath := range iterationPaths {
			if iterationPath.distance < shortestPathIteration.distance {
				shortestPathIteration = iterationPath
			}
		}
		if shortestPathIteration.distance < shortestPathAllTime.distance {
			shortestPathAllTime = shortestPathIteration
		}
		//fmt.Println("shortestPathIteration", shortestPathIteration)
		//fmt.Println("shortestPathAllTime", shortestPathAllTime)
	}
	return shortestPathAllTime
}

func (a *AntColony) genPaths() []TourPath {
	var iterationPaths []TourPath
	for ant := 0; ant < a.noAnts; ant++ {
		start := 0
		iterationPath := a.genPath(start)
		iterationPaths = append(iterationPaths, TourPath{iterationPath, a.pathDistance(iterationPath)})
	}
	return iterationPaths
}

func (a *AntColony) genPath(start int) []unitPath {
	var antPath []unitPath
	visitedNodes := set.New(set.ThreadSafe)
	visitedNodes.Add(start)
	//start := start
	//fmt.Println(pheromones)
	for i := 0; i < len(a.distances)-1; i++ {
		end := a.pickNextMove(pheromones[start], a.distances[start], visitedNodes.List())
		antPath = append(antPath, unitPath{start, end})
		start = end
		visitedNodes.Add(end)
	}
	return antPath
}

func (a *AntColony) pickNextMove(pheromone []float64, distances []float64, visited []interface{}) int {
	//pheromone := pheromone
	for i := range visited {
		pheromone[i] = 0
	}
	var row []float64
	//distances := distances
	for i := 0; i < len(pheromone); i++ {
		row = append(row, math.Pow(pheromone[i], a.alpha)*(math.Pow((1.0/distances[i]), a.beta)))
	}

	rowSum := 0
	for i := range row {
		rowSum += i
	}
	for i := 0; i < len(row); i++ {
		row[i] = row[i] / float64(rowSum)
	}
	choices := make([]weightedrandom.Choice, 0, len(pheromone))
	for i := 0; i < len(pheromone); i++ {
		choices = append(choices, weightedrandom.Choice{row[i], a.nodes[i]})
	}
	choice := weightedrandom.WeightedChoice(choices).Item
	node, _ := choice.(int)

	return node
}

func (a *AntColony) pathDistance(iterationPath []unitPath) float64 {
	distance := 0.0
	for _, edge := range iterationPath {
		distance += a.distances[edge.start][edge.end]
	}
	return distance
}

func (a *AntColony) spreadPheromone(iterationPaths []TourPath, noBestPath int) {
	//bestPathsIteration := sort(iterationPaths, noBestPath)
	sort.Slice(iterationPaths, func(i, j int) bool { return iterationPaths[i].distance < iterationPaths[j].distance })
	if noBestPath > a.noAnts {
		noBestPath = a.noAnts
	}
	for _, path := range iterationPaths[0:noBestPath] {
		for _, edge := range path.path {
			pheromones[edge.start][edge.end] += 1.0 / (a.distances[edge.start][edge.end])
		}
	}
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
			dis, _ := strconv.ParseFloat(str, 10)
			//fmt.Println(line[i])
			if dis == 0.0 {
				row = append(row, MaxInt)
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

	var antColony AntColony
	antColony = AntColony{
		noAnts:      500,
		noBestPath:  450,
		noIteration: 200,
		decay:       0.85,
		alpha:       1,
		beta:        1,
		distances:   distances,
	}

	shortestPath := antColony.Run()
	fmt.Println(shortestPath)

}
