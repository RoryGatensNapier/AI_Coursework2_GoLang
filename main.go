package main

import (
	"fmt"

	FH "github.com/RoryGatensNapier/AI_Coursework2_GoLang/FileHandler"
)

func main() {
	//enter main code here from resources functions
	data := FH.ReadFromFile("./cavernFiles/input1.cav")
	FH.ConvertElementsToInt(data)
	fmt.Println(data)
	//fmt.Println(m.Construct(data))
}
