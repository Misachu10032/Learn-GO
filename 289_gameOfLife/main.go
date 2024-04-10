package main

import "fmt"

func main() {
	
	// aa := minWindow(s, t)
	// fmt.Print(aa)
	array := [][]int{
        {0, 1, 0},
        {0, 0, 1},
        {1, 1, 1},
        {0, 0, 0},
    }

	gameOfLife(array)

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func checkYourNeighbour(board [][]int, row, col int) int {
	tempCount := 0
	maxCol := len(board[0])-1
	maxRow := len(board)-1
	for i := max(0, row-1); i <= min(row+1, maxRow); i++ {
		for j := max(0, col-1); j <= min(col+1, maxCol); j++ {

			if i == row && j == col {
				continue
			}

			tempCount += board[i][j]
		}
	}
	return tempCount
}

func gameOfLife(board [][]int) {

	col := len(board[0])
	row := len(board)
	nextState := make([][]int, row)
	for i := range nextState {
		nextState[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {

			count := checkYourNeighbour(board, i, j)
             fmt.Print(count)
             fmt.Print("\n")
			if board[i][j]==1 {
				if count < 2 {
					nextState[i][j] = 0
				}
				if count > 3 {
					nextState[i][j] = 0
				}
				if  count >=2 && count<= 3 {
					nextState[i][j] = 1
				}
			} else {
				if count == 3 {
					nextState[i][j] = 1
				} else {
					nextState[i][j] = 0
				}
			}

		}
	}


    	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			board[i][j] = nextState[i][j]
		}
	}
}
