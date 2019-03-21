package acopar

import (
	"experiments/ant_colony_optimization/weightedrandom"
	"math"
	"sort"

	"github.com/fatih/set"
)

//Ant colony struct
type AntColony struct {
	NoAnts      int
	NoBestPath  int
	NoIteration int
	Decay       float64
	Alpha       float64
	Beta        float64
	Distances   [][]float64
	Nodes       []int
}

type unitPath struct {
	start int
	end   int
}

var (
	pheromones [][]float64
)

type TourPath struct {
	Path     []unitPath
	Distance float64
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = float64(int(MaxUint >> 1))

//Run() performs the ant colony optimization and returns the shortest tour path among all iterations.
func (a *AntColony) Run() TourPath {
	//fmt.Println(MaxInt)
	for i := 0; i < len(a.Distances); i++ {
		a.Nodes = append(a.Nodes, i)
	}

	for row := 0; row < len(a.Distances); row++ {
		var rowValues []float64
		for value := 0; value < len(a.Distances[row]); value++ {
			rowValues = append(rowValues, 1.0/float64(len(a.Distances)))
		}
		pheromones = append(pheromones, rowValues)
	}

	var shortestPathIteration TourPath
	var shortestPathAllTime TourPath
	shortestPathAllTime.Distance = MaxInt
	for iteration := 0; iteration < a.NoIteration; iteration++ {
		var iterationPaths []TourPath
		iterationPaths = a.genPaths()
		a.spreadPheromone(iterationPaths, a.NoBestPath)
		//fmt.Println("pheromones----->", pheromones)
		shortestPathIteration = iterationPaths[0]
		for _, iterationPath := range iterationPaths {
			if iterationPath.Distance < shortestPathIteration.Distance {
				shortestPathIteration = iterationPath
			}
		}
		if shortestPathIteration.Distance < shortestPathAllTime.Distance {
			shortestPathAllTime = shortestPathIteration
		}
		//fmt.Println("shortestPathIteration", shortestPathIteration)
		//fmt.Println("shortestPathAllTime", shortestPathAllTime)
	}
	return shortestPathAllTime
}

func (a *AntColony) genPaths() []TourPath {
	var iterationPaths []TourPath
	//fmt.Println("pheromones---->", pheromones)
	iterationPathChan := make(chan []unitPath, a.NoAnts*100)
	for ant := 0; ant < a.NoAnts; ant++ {
		start := 0
		go a.genPath(start, iterationPathChan, ant)
	}
	for i := 0; i < a.NoAnts; i++ {
		//fmt.Println("getting from ant ", i)
		iterationPath := <-iterationPathChan
		iterationPaths = append(iterationPaths, TourPath{iterationPath, a.pathDistance(iterationPath)})
	}
	return iterationPaths
}

func (a *AntColony) genPath(start int, iterationPathChan chan []unitPath, ant int) {
	//fmt.Println("going inside ", ant)
	var antPath []unitPath
	visitedNodes := set.New(set.ThreadSafe)
	visitedNodes.Add(start)
	//start := start
	prev := start
	//fmt.Println(prev)
	//fmt.Println(pheromones)
	for i := 0; i < len(a.Distances)-1; i++ {
		pheromonePrev := make([]float64, len(pheromones[prev]))
		copy(pheromonePrev, pheromones[prev])
		next := a.pickNextMove(pheromonePrev, a.Distances[prev], visitedNodes.List())
		//fmt.Println(next)
		antPath = append(antPath, unitPath{prev, next})
		//fmt.Println(prev, next)
		prev = next
		visitedNodes.Add(next)
	}
	//fmt.Println("visitedNodes", visitedNodes.Size(), visitedNodes.List())
	antPath = append(antPath, unitPath{prev, start})
	//fmt.Println("antPath", antPath)
	//fmt.Println("pheromones---->", pheromones)

	//fmt.Println("getting out of ant ", ant)
	iterationPathChan <- antPath
	//return antPath
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func (a *AntColony) pickNextMove(pheromone []float64, Distances []float64, visited []interface{}) int {
	//pheromone := pheromone
	//fmt.Println("visited", visited)
	//fmt.Println("pheromone", pheromone)
	for i := 0; i < len(visited); i++ {
		value, _ := visited[i].(int)
		//fmt.Println(value)
		//pheromone[visited[i]] = 0
		pheromone[value] = 0
	}
	//for i := range visited {
	//	pheromone[i] = 0
	//}
	//fmt.Println("pheromone", pheromone)
	var row []float64
	//Distances := Distances

	//fmt.Println(pheromone, Distances)
	for i := 0; i < len(pheromone); i++ {
		value := math.Pow(pheromone[i], a.Alpha) * (math.Pow((1.0 / Distances[i]), a.Beta))
		//value = float64(int(value*100000000000000)) / 100000000000000
		//value = toFixed(value, 14)
		row = append(row, value)
	}

	rowSum := 0.0
	for _, i := range row {
		rowSum += i
	}
	//fmt.Println("row", row)
	//fmt.Println("rowSum", rowSum)
	for i := 0; i < len(row); i++ {
		value := row[i] / float64(rowSum)
		//value = toFixed(value, 8)
		row[i] = value
	}
	//fmt.Println("rownorm", row)
	choices := make([]weightedrandom.Choice, 0, len(pheromone))
	for i := 0; i < len(pheromone); i++ {
		choices = append(choices, weightedrandom.Choice{row[i], a.Nodes[i]})
	}
	choice := weightedrandom.WeightedChoice(choices).Item
	node, _ := choice.(int)

	return node
}

func (a *AntColony) pathDistance(iterationPath []unitPath) float64 {
	distance := 0.0
	for _, edge := range iterationPath {
		distance += a.Distances[edge.start][edge.end]
	}
	return distance
}

func (a *AntColony) spreadPheromone(iterationPaths []TourPath, NoBestPath int) {
	//bestPathsIteration := sort(iterationPaths, NoBestPath)
	sort.Slice(iterationPaths, func(i, j int) bool { return iterationPaths[i].Distance < iterationPaths[j].Distance })
	if NoBestPath > a.NoAnts {
		NoBestPath = a.NoAnts
	}
	for _, path := range iterationPaths[0:NoBestPath] {
		for _, edge := range path.Path {
			//fmt.Println(edge.start, edge.end, pheromones[edge.start][edge.end])
			pheromones[edge.start][edge.end] = pheromones[edge.start][edge.end] + 1.0/(a.Distances[edge.start][edge.end])
			//fmt.Println(edge.start, edge.end, pheromones[edge.start][edge.end])
		}
	}
}
