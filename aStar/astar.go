package astar

import (
	"fmt"
	"math"
	"sync"

	dt "github.com/RoryGatensNapier/AI_Coursework2_GoLang/DataTypes"
)

func constructCave(id int, pos dt.Vec2, truths []int) dt.Cave {
	NewCave := dt.Cave{ID: id, Position: pos, ConnectsTo: truths, HeuristicScore: -1}
	return NewCave
}

func ConstructCaves(CavSys dt.CavernSystem) []dt.Cave {
	var allCaves []dt.Cave
	var conv_truths []int
	for i := 0; i < len(CavSys.Points); i++ {
		for _, v := range CavSys.Truths {
			conv_truths = append(conv_truths, v[i])
		}
		newCave := constructCave(i+1, CavSys.Points[i+1], conv_truths)
		allCaves = append(allCaves, newCave)
		conv_truths = nil
	}
	return allCaves
}

func EuclideanDistance(curCave dt.Cave, goalCave dt.Cave) float64 {
	dist := math.Sqrt(math.Pow(float64(goalCave.Position.X-curCave.Position.X), 2) + math.Pow(float64(goalCave.Position.Y-curCave.Position.Y), 2))
	return dist
}

func ScoreWorker(curCave dt.Cave, neighbour dt.Cave, endCave dt.Cave, result chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	n_dist := EuclideanDistance(curCave, neighbour)
	e_dist := EuclideanDistance(neighbour, endCave)
	score := n_dist + e_dist
	result <- score
}

func ExpandNode(ToExpand dt.Cave, Caves []dt.Cave) {
	var connectedNodes []int
	wg := &sync.WaitGroup{}
	for i, v := range ToExpand.ConnectsTo {
		if v == 1 {
			connectedNodes = append(connectedNodes, i+1)
		}
	}
	scoresChan := make(chan float64, len(connectedNodes))
	fmt.Println(connectedNodes)
	wg.Add(len(connectedNodes))
	for _, v := range connectedNodes {
		//go EuclideanDistance(Caves[v], Caves[len(Caves)-1])
		go ScoreWorker(ToExpand, Caves[v-1], Caves[len(Caves)-1], scoresChan, wg)
		fmt.Println("making it")
	}
	go func() {
		wg.Wait()
		close(scoresChan)
	}()
	for x := range scoresChan {
		fmt.Println(x)
	}
	fmt.Println("made it")
}

func ConstructPath(path []dt.Cave, newCave dt.Cave) {

}

func DoAStar(CavSys dt.CavernSystem) {
	Caves := ConstructCaves(CavSys)

	//involve min heap where heap is open nodes

	start := Caves[0]
	end := Caves[len(Caves)-1]
	initHeuristic := EuclideanDistance(start, end)
	Caves[0].HeuristicScore = initHeuristic
	var openCaves dt.CaveHeap = dt.CaveHeap{&Caves[0]}
	var visitedCaves []dt.Cave
	openCaves.Init()
	fmt.Println(*openCaves[0])
	fmt.Println(end)
	for len(openCaves) > 0 {
		currentCave := *openCaves[0]
		// if currentCave.ID == end.ID {
		// 	//return foundPath

		visitedCaves = append(visitedCaves, *openCaves[0])
		openCaves.Pop()
		ExpandNode(currentCave, Caves)
	}
	//visited.append(Caves[0])
}
