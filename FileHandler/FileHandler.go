package FileHandler

import (
	"encoding/csv"
	//"fmt"
	"log"
	"os"
	"strconv"
)

func ReadFromFile(filePath string) []string {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
		file.Close()
		return nil
	}
	readCav := csv.NewReader(file)
	vals, err := readCav.Read()
	if err != nil {
		log.Fatal("Could not parse .CAV file at "+filePath, err)
	}
	file.Close()
	return vals
}

func ConvertElementsToInt(FileData []string) []int {
	var retVal []int
	for _, val_F := range FileData {
		result, err := strconv.Atoi(val_F)
		if err != nil {
			log.Fatal("Issue in converting values to int!", err)
		}
		retVal = append(retVal, result)
	}
	return retVal
}
