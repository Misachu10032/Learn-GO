package main

import (
	"fmt"
)

func main() {

	// aa := minWindow(s, t)
	// fmt.Print(aa)
	strs := [][]int{{1, 3}, {6, 9}}
	aa := []int{2, 5}
	fmt.Print(insert(strs, aa))

}

func greater(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func lesser(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	i:=0
	for i < len(intervals) && intervals[i][1] < newInterval[0]  {
		result = append(result, intervals[i])
		i++
	}

	fmt.Print(result)
	fmt.Print("\n")
	temp:=[]int{newInterval[0],newInterval[1]}
	for  i < len(intervals) && intervals[i][1] > newInterval[0] && intervals[i][0] < newInterval[1]{
		temp[0] = lesser(intervals[i][0], temp[0])
		temp[1] =  greater(intervals[i][1], temp[1])
		i++
	}


	result = append(result, temp)


	fmt.Print(result)
	fmt.Print("\n")
	for i < len(intervals)   {

		fmt.Print("neverere")
		result = append(result, intervals[i])
		i++
	}

	return result
}
