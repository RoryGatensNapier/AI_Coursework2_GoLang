package DataTypes

type Vec2 struct {
	X, Y int
}

type Cave struct {
	ID             int
	Position       Vec2
	ConnectsTo     []int
	HeuristicScore float64
}

type CavernSystem struct {
	Points map[int]Vec2
	Truths [][]int
}
