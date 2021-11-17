package astar

import (
	dt "github.com/RoryGatensNapier/AI_Coursework2_GoLang/DataTypes"
)

type Vec2 = dt.Vec2
type Cave = dt.Cave

func ConstructCave(id int, vecMap map[int]Vec2, truthMatrix [][]int) Cave {
	NewCave := Cave{ID: id}
	NewCave.Position = vecMap[id]
	NewCave.ConnectsTo = truthMatrix[id]
	return NewCave
}
