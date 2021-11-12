package main

import (
	"fmt"

	FileHandler "github.com/RoryGatensNapier/AI_Coursework2_GoLang/FileHandler"
	matrix "github.com/RoryGatensNapier/AI_Coursework2_GoLang/Math"
)

func main() {
	//enter main code here from resources functions
	data := FileHandler.ReadFromFile("./cavernFiles/input1.cav")
	fmt.Println(matrix.Construct(data))
}
