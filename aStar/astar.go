package astar

import (
	"container/heap"
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

func ScoreWorker(curCave dt.Cave, neighbour dt.Cave, endCave dt.Cave, result chan<- dt.Cave, wg *sync.WaitGroup) {
	defer wg.Done()
	n_dist := EuclideanDistance(curCave, neighbour) + curCave.TraversedDistance
	e_dist := EuclideanDistance(neighbour, endCave)
	score := n_dist + e_dist
	neighbour.TraversedDistance = n_dist
	neighbour.HeuristicScore = score
	result <- neighbour
}

func ExpandNode(ToExpand dt.Cave, Caves []dt.Cave, visitedCaves []dt.Cave) []dt.Cave {
	var connectedNodes []int
	var resultNodes []dt.Cave
	wg := &sync.WaitGroup{}
	for i, v := range ToExpand.ConnectsTo {
		if v == 1 {
			if len(visitedCaves) > 0 {
				for _, vis := range visitedCaves {
					if vis.ID != i+1 {
						connectedNodes = append(connectedNodes, i+1)
						break
					}
				}
			} else {
				connectedNodes = append(connectedNodes, i+1)
			}
		}
	}
	fmt.Println("Values in connectedNodes = ", connectedNodes)
	scoresChan := make(chan dt.Cave, len(connectedNodes))
	wg.Add(len(connectedNodes))
	for _, v := range connectedNodes {
		go ScoreWorker(ToExpand, Caves[v-1], Caves[len(Caves)-1], scoresChan, wg)
	}
	wg.Wait()
	close(scoresChan)
	for x := range scoresChan {
		fmt.Println(x, " result added")
		resultNodes = append(resultNodes, x)
	}
	fmt.Println("made it")
	return resultNodes
}

func findNeighbours(ToExpand dt.Cave, Caves []dt.Cave, visitedCaves []dt.Cave) []int {
	var connectedNodes []int
	for i, v := range ToExpand.ConnectsTo {
		if v == 1 {
			if len(visitedCaves) > 0 {
				for _, vis := range visitedCaves {
					if vis.ID != i+1 {
						connectedNodes = append(connectedNodes, i+1)
						break
					}
				}
			} else {
				connectedNodes = append(connectedNodes, i+1)
			}
		}
	}
	fmt.Println("Values in connectedNodes = ", connectedNodes)
	return connectedNodes
}

func calcScores(nieghbour_ids []int, ToExpand dt.Cave, Caves []dt.Cave, visitedCaves []dt.Cave) []dt.Cave {
	var resultingCaves []dt.Cave
	scoresChan := make(chan dt.Cave, len(nieghbour_ids))
	wg := &sync.WaitGroup{}
	wg.Add(len(nieghbour_ids))
	for _, v := range nieghbour_ids {
		go ScoreWorker(ToExpand, Caves[v-1], Caves[len(Caves)-1], scoresChan, wg)
	}
	wg.Wait()
	close(scoresChan)
	for x := range scoresChan {
		resultingCaves = append(resultingCaves, x)
	}
	fmt.Println("made it")
	return resultingCaves
}

func DoAStar(CavSys dt.CavernSystem) {
	Caves := ConstructCaves(CavSys)
	start := Caves[0]
	end := Caves[len(Caves)-1]
	initHeuristic := EuclideanDistance(start, end)
	Caves[0].HeuristicScore = initHeuristic
	var openCaves dt.CaveHeap = dt.CaveHeap{Caves[0]}
	var visitedCaves []dt.Cave
	openCaves.Init()
	for len(openCaves) > 0 {
		currentCave := openCaves[0]
		fmt.Println("Current active cave is ", currentCave)
		fmt.Println("Visited caves = ", visitedCaves)
		if currentCave.ID == end.ID {
			//return foundPath
			fmt.Println("Found end!")
			return
		}
		for _, v := range visitedCaves {
			if v.ID == currentCave.ID {
				fmt.Println("already visited this node!")
				fmt.Println("current cave: ", currentCave.ID)
				fmt.Println("test cave: ", v.ID)
				for test_i, test_v := range openCaves {
					fmt.Println("Cave ", test_i, " is ", test_v)
				}
				return
			}
		}
		visitedCaves = append(visitedCaves, openCaves[0])
		//openCaves.Pop()
		openCaves.Remove(0)
		openCaves.Fix(0)
		newNodes := ExpandNode(currentCave, Caves, visitedCaves)
		for _, v := range newNodes {
			heap.Push(&openCaves, v)
		}
		//openCaves.Init()
		fmt.Println("top cave now equals = ", openCaves[0])
		fmt.Println("run success")
	}
	//visited.append(Caves[0])
}

func DoAStar_V2(CavSys dt.CavernSystem) {
	Caves := ConstructCaves(CavSys)
	start := Caves[0]
	end := Caves[len(Caves)-1]
	initHeuristic := EuclideanDistance(start, end)
	Caves[0].HeuristicScore = initHeuristic
	var openCaves dt.CaveHeap = dt.CaveHeap{Caves[0]}
	var visitedCaves []dt.Cave = []dt.Cave{openCaves[0]}
	openCaves.Init()
	for len(openCaves) > 0 {
		for _, op_val := range openCaves {
			if op_val.ID == end.ID {
				//return foundPath
				fmt.Println("Found end!")
				return
			}
			for _, v := range visitedCaves {
				if v.ID == op_val.ID {
					fmt.Println("already visited this node!")
					fmt.Println("current cave: ", op_val)
					fmt.Println("test cave: ", v)
					for test_i, test_v := range openCaves {
						fmt.Println("Cave ", test_i, " is ", test_v)
					}
					return
				}
			}
			exploreNeighbours := findNeighbours(op_val, Caves, visitedCaves)
			scoredNeighbours := calcScores(exploreNeighbours, op_val, Caves, visitedCaves)
		}
	}
}
