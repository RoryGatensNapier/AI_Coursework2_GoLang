package astar

import (
	//"fmt"
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

func NC_ScoreWorker(curCave dt.Cave, neighbour dt.Cave, endCave dt.Cave) dt.Cave {
	n_dist := EuclideanDistance(curCave, neighbour) + curCave.TraversedDistance
	e_dist := EuclideanDistance(neighbour, endCave)
	score := n_dist + e_dist
	neighbour.TraversedDistance = n_dist
	neighbour.HeuristicScore = score
	return neighbour
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
					} else if vis.ID == i+1 && Caves[i].HeuristicScore < vis.HeuristicScore {
						//fmt.Println("found node through lesser cost?")
					}
				}
			} else {
				connectedNodes = append(connectedNodes, i+1)
			}
		}
	}
	//fmt.Println("Values in connectedNodes = ", connectedNodes)
	return connectedNodes
}

func calcScores(nieghbour_ids []int, ToExpand dt.Cave, Caves []dt.Cave, visitedCaves []dt.Cave) []dt.Cave {
	var resultingCaves dt.CaveHeap
	scoresChan := make(chan dt.Cave, len(nieghbour_ids))
	wg := &sync.WaitGroup{}
	wg.Add(len(nieghbour_ids))
	for _, v := range nieghbour_ids {
		go ScoreWorker(ToExpand, Caves[v-1], Caves[len(Caves)-1], scoresChan, wg)
	}
	wg.Wait()
	close(scoresChan)
	for x := range scoresChan {
		for _, v := range ToExpand.FoundFrom {
			if ToExpand.ID == v {
				continue
			}
		}
		foundFromConstruct := append(ToExpand.FoundFrom, ToExpand.ID)
		x.FoundFrom = append(x.FoundFrom, foundFromConstruct...)
		resultingCaves = append(resultingCaves, x)
	}
	//fmt.Println("made it")
	return resultingCaves
}

func NC_calcScores(nieghbour_ids []int, ToExpand dt.Cave, Caves []dt.Cave, visitedCaves []dt.Cave) []dt.Cave {
	var resultingCaves dt.CaveHeap
	var pre_scores []dt.Cave
	for _, v := range nieghbour_ids {
		pre_scores = append(pre_scores, NC_ScoreWorker(ToExpand, Caves[v-1], Caves[len(Caves)-1]))
	}
	for _, x := range pre_scores {
		foundFromConstruct := append(ToExpand.FoundFrom, ToExpand.ID)
		x.FoundFrom = append(x.FoundFrom, foundFromConstruct...)
		resultingCaves = append(resultingCaves, x)
	}
	//fmt.Println("made it")
	return resultingCaves
}

func DoAStar_V2(CavSys dt.CavernSystem) []int {
	Caves := ConstructCaves(CavSys)
	start := Caves[0]
	end := Caves[len(Caves)-1]
	initHeuristic := EuclideanDistance(start, end)
	Caves[0].HeuristicScore = initHeuristic
	var openCaves dt.CaveHeap = dt.CaveHeap{Caves[0]}
	var visitedCaves []dt.Cave
	var path []int
	openCaves.Init()
	for len(openCaves) > 0 {
		currentCave := &openCaves[0]
		//fmt.Println("current cave is ", *currentCave)
		if currentCave.ID == end.ID {
			for _, v := range visitedCaves {
				if v.ID == currentCave.ID && v.HeuristicScore > currentCave.HeuristicScore {
					*currentCave = v
				}
			}
			currentCave.FoundFrom = append(currentCave.FoundFrom, currentCave.ID)
			//fmt.Println(openCaves)
			fmt.Println(currentCave.FoundFrom)
			//fmt.Println(currentCave.TraversedDistance)
			return path
		}
		for i, v := range visitedCaves {
			if v.ID == currentCave.ID {
				if currentCave.HeuristicScore < v.HeuristicScore {
					//fmt.Println(openCaves)
					visitedCaves[i] = *currentCave
					//fmt.Println(visitedCaves)
					break
				} else {
					*currentCave = openCaves[1]
				}
			}
		}
		exploreNeighbours := findNeighbours(*currentCave, Caves, visitedCaves)
		scoredNeighbours := calcScores(exploreNeighbours, *currentCave, Caves, visitedCaves)
		visitedCaves = append(visitedCaves, *currentCave)
		openCaves.Remove(0)
		for _, v := range scoredNeighbours {
			openCaves.Push(v)
		}
		// fmt.Println("Caves = ", openCaves)
		// fmt.Println("visited caves = ", visitedCaves)
	}
	return nil
}
