package FileHandler

import (
	"log"
	"os"
	"time"
)

func GiveTime() time.Time {
	return time.Now()
}

func OpenFile_Read(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
		file.Close()
		return nil
	}
	file.Close()
	return file
}
