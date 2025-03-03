package l

import (
	"awesomeProject/matrix"
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {

	var a = 5
	var b = 6

	// Printing multiplication
	fmt.Printf("%d * %d = %d\n", a, b, a*b)

	// Add function example
	result := add(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, result)

	// Declare and print array
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 2D array
	twoD := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(twoD)
	vertex1 := matrix.Vertex{1, 2}
	t := vertex1.Sq(1, 2)
	fmt.Println("vertex squared value ", t)
	p := vertex1.Sqrt(t)
	fmt.Println("vertex squared value ", p)
	// Create matrices and perform operations
	matrix1 := matrix.NewMatrix(2, 3)
	matrix1.SetElement(0, 0, 1)
	matrix1.SetElement(0, 1, 2)
	matrix1.SetElement(0, 2, 3)
	matrix1.SetElement(1, 0, 4)
	matrix1.SetElement(1, 1, 5)
	matrix1.SetElement(1, 2, 6)

	matrix2 := matrix.NewMatrix(2, 3)
	matrix2.SetElement(0, 0, 7)
	matrix2.SetElement(0, 1, 8)
	matrix2.SetElement(0, 2, 9)
	matrix2.SetElement(1, 0, 10)
	matrix2.SetElement(1, 1, 11)
	matrix2.SetElement(1, 2, 12)

	// Print matrices
	fmt.Println("Matrix 1:")
	matrix1.PrintMatrix()
	fmt.Println("Matrix 2:")
	matrix2.PrintMatrix()

	// Add matrices
	// Correctly assign the result to a Matrix, not an int
	matrixResult, err := matrix1.Add(matrix2)
	if err != nil {
		fmt.Println("Error adding matrices:", err)
		return
	}
	fmt.Println("Matrix Addition Result:")

	matrixResult.PrintMatrix()
}
