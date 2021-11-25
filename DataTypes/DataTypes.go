package DataTypes

import (
	"container/heap"
	//"sort"
)

type Vec2 struct {
	X, Y int
}

type Cave struct {
	ID                int
	Position          Vec2
	ConnectsTo        []int
	TraversedDistance float64
	HeuristicScore    float64
}

type CaveHeap []Cave

func (cc CaveHeap) Len() int { return len(cc) }
func (cc CaveHeap) Less(idx_1, idx_2 int) bool {
	return cc[idx_1].HeuristicScore < cc[idx_2].HeuristicScore
}
func (cc CaveHeap) Swap(idx_1, idx_2 int) { cc[idx_1], cc[idx_2] = cc[idx_2], cc[idx_1] }

func (cc *CaveHeap) Push(x interface{}) {
	*cc = append(*cc, x.(Cave))
}

func (cc *CaveHeap) Pop() interface{} {
	old := *cc
	n := len(old)
	retCave := old[n-1]
	*cc = old[0 : n-1]
	return retCave
}

func (cc *CaveHeap) Remove(i int) {
	heap.Remove(cc, i)
}

func (cc *CaveHeap) Fix(i int) {
	heap.Fix(cc, i)
}

func (cc CaveHeap) Init() {
	heap.Init(&cc)
}

type CavernSystem struct {
	Points map[int]Vec2
	Truths [][]int
}
