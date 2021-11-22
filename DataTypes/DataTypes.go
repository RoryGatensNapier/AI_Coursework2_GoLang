package DataTypes

import (
	"container/heap"
	//"sort"
)

type Vec2 struct {
	X, Y int
}

type Cave struct {
	ID             int
	HeapID         int
	Position       Vec2
	ConnectsTo     []int
	HeuristicScore float64
}

type CaveHeap []*Cave

func (cc CaveHeap) Len() int { return len(cc) }
func (cc CaveHeap) Less(idx_1, idx_2 int) bool {
	return cc[idx_1].HeuristicScore < cc[idx_2].HeuristicScore
}
func (cc CaveHeap) Swap(idx_1, idx_2 int) { cc[idx_1], cc[idx_2] = cc[idx_2], cc[idx_1] }

func (cc *CaveHeap) Push(x interface{}) {
	n := len(*cc)
	pushCave := x.(*Cave)
	pushCave.HeapID = n
	*cc = append(*cc, pushCave)
}

func (cc *CaveHeap) Pop() interface{} {
	old := *cc
	n := len(old)
	retCave := old[n-1]
	old[n-1] = nil
	*cc = old[0 : n-1]
	return retCave
}

func (cc *CaveHeap) update(CaveToUpdate *Cave, new_id int, new_pos Vec2, new_con []int, new_score float64, new_HeapID int) {
	CaveToUpdate.ID = new_id
	CaveToUpdate.Position = new_pos
	CaveToUpdate.ConnectsTo = new_con
	CaveToUpdate.HeuristicScore = new_score
	CaveToUpdate.HeapID = new_HeapID
	heap.Fix(cc, CaveToUpdate.ID)
}

func (cc CaveHeap) Init() {
	heap.Init(&cc)
}

type CavernSystem struct {
	Points map[int]Vec2
	Truths [][]int
}
