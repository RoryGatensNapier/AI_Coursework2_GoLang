package main

import (
	"fmt"

	FH "github.com/RoryGatensNapier/AI_Coursework2_GoLang/FileHandler"
	matrices "github.com/RoryGatensNapier/AI_Coursework2_GoLang/Math"
)

func main() {
	//enter main code here from resources functions
	data := FH.ReadFromFile("./cavernFiles/input1.cav")
	new_ints := FH.ConvertElementsToInt(data)
	new_matrix := matrices.ConstructTruthTable(new_ints)
	new_vecMap := matrices.ConstructVectorMapping(new_ints)
	fmt.Println(new_matrix)
	fmt.Println(new_vecMap)
	//fmt.Println(m.Construct(data))
}
