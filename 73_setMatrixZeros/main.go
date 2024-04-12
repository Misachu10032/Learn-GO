package main

import "fmt"

func main() {

	// aa := minWindow(s, t)
	// fmt.Print(aa)
	array := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	b := array
	fmt.Print(b)
	setZeroes(array)

}
func setZeroes(matrix [][]int) {
	rows := len(matrix)
    cols := len(matrix[0])
    resultMatrix := make([][]int, rows)
    for i := range matrix {
        resultMatrix[i] = make([]int, cols)
        copy(resultMatrix[i], matrix[i])
    }

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {

			if matrix[i][j] == 0 {
	
				for k := 0; k < len(matrix); k++ {
					resultMatrix[k][j] = 0
				}
				for k := 0; k < len(matrix[0]); k++ {
					resultMatrix[i][k] = 0
				}
			}
		}
	}


        	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] = resultMatrix[i][j]
		}
	}
}