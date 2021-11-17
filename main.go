package main

import (
	"fmt"
	//"os"

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
	new_matrix := matrices.ConstructTruthTable(new_ints)
	new_vecMap := matrices.ConstructVectorMapping(new_ints)
	//fmt.Println("Truth Matrix = ", new_matrix)
	fmt.Println("Vector Map = ", new_vecMap)
	start := astar.ConstructCave(1, new_vecMap, new_matrix)
	fmt.Println(start)
	//fmt.Println(m.Construct(data))
}
