package matrix

//import statments
import (
	"fmt"
	"log"
)

// "log"

type Vec2 struct {
	x, y int
}

func ConstructTruthTable(FileData []int) [][]int {
	if len(FileData) == 0 {
		log.Fatalln("No data passed into ConstructFileData!", log.Llongfile)
	}
	dimension := FileData[0]
	locations := FileData[1 : dimension*2+1]
	truths := FileData[dimension*2+1:]
	fmt.Println("Raw Truths = ", truths)
	fmt.Println("Raw Locations = ", locations)

	matrix := make([][]int, dimension)
	for i := 0; i < dimension; i++ {
		matrix[i] = truths[dimension*i : dimension+dimension*i] //make([]int, dimension)
	}
	return matrix
}

func ConstructVectorMapping(FileData []int) map[int]Vec2 {
	dimension := FileData[0]
	locations := FileData[1 : dimension*2+1]

	vectorMappings := make(map[int]Vec2)
	for i := 0; i < dimension*2-1; i++ {
		vecSlice := locations[i : i+2]
		var (
			x, y int
		)
		if i%2 != 0 {
			continue
		} else {
			x = vecSlice[0]
			y = vecSlice[1]
			//_ = x + y
			vectorMappings[(i/2)+1] = Vec2{x, y}
		}
	}
	return vectorMappings
}
