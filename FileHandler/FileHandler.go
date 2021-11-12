package FileHandler

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func GiveTime() time.Time {
	return time.Now()
}

func ReadFromFile(filePath string) [][]string {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal("Could not open file at "+filePath, err)
		file.Close()
		return nil
	}
	readCav := csv.NewReader(file)
	vals, err := readCav.ReadAll()
	if err != nil {
		log.Fatal("Could not parse .CAV file at "+filePath, err)
	}
	//file.Close()
	fmt.Println(vals)
	return vals
}
