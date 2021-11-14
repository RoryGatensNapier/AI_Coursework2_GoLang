package matrix

//import statments
import "fmt"

// "log"

type Vec2 struct {
	x, y int
}

func ConstructTruthTable(FileData []int) [][]int {
	dimension := FileData[0]
	locations := FileData[1 : dimension*2+1]
	truths := FileData[dimension*2+1:]
	fmt.Println("Truths = ", truths)
	fmt.Println("Locations = ", locations)

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
			vectorMappings[i] = Vec2{x, y}
		}
	}
	return vectorMappings
}
