package FileHandler

import (
	"log"
	"os"
	"time"
)

func GiveTime() time.Time {
	return time.Now()
}

func ReadFromFile(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
		file.Close()
		return nil
	}
	file.Close()
	return file
}
