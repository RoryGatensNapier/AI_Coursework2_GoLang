package matrix

//import statments
import "fmt"

// "log"

type Vec2 struct {
	x, y float64
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
