package matrix

//import statments
import "fmt"

// "log"
// "strconv"

type Vec2 struct {
	x float64
	y float64
}

func ConstructTruthTable(FileData []int) [][]int {
	dimension := FileData[0]
	matrix := make([][]int, dimension)
	for i := 0; i < dimension; i++ {
		matrix[i] = make([]int, dimension)
	}
	locations := FileData[1 : dimension*2]
	truths := FileData[dimension*2:]
	fmt.Println("Truths = ", truths)
	fmt.Println("Locations = ", locations)
	return matrix
}
