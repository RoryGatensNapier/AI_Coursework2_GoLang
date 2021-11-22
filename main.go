package main

import (
	//"fmt"
	//"os"
	dt "github.com/RoryGatensNapier/AI_Coursework2_GoLang/DataTypes"
	FH "github.com/RoryGatensNapier/AI_Coursework2_GoLang/FileHandler"
	matrices "github.com/RoryGatensNapier/AI_Coursework2_GoLang/Math"
	astar "github.com/RoryGatensNapier/AI_Coursework2_GoLang/aStar"
)

func main() {
	//os_args := os.Args[1:]
	//fmt.Println(os_args[0])
	//enter main code here from resources functions
	//data := FH.ReadFromFile("./" + os_args[0] + ".cav")
	data := FH.ReadFromFile("./cavernFiles/input1.cav")
	new_ints := FH.ConvertElementsToInt(data)
	CavSys := dt.CavernSystem{Points: matrices.ConstructVectorMapping(new_ints), Truths: matrices.ConstructTruthTable(new_ints)}
	//fmt.Println("Truth Matrix = ", new_matrix)
	//start := astar.ConstructCave(1, new_vecMap, new_matrix)
	//fmt.Println(CavSys)
	astar.DoAStar(CavSys)
}
