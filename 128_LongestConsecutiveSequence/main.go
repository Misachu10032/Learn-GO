package main

import (
	"fmt"
	"sort"
)

func main() {

	// aa := minWindow(s, t)
	// fmt.Print(aa)
	strs := []int{100, 4, 200, 1, 3, 2}
	fmt.Print(longestConsecutive(strs))

}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	array := nums
	sort.Ints(array)
	result := 1
	temp := 1
	duplicateRemoved := make([]int, 0)

	for i := 1; i < len(array); i++ {

		if array[i] != array[i-1] {
			duplicateRemoved = append(duplicateRemoved, array[i])
		}
	}
    fmt.Print(duplicateRemoved)
	for i := 0; i < len(duplicateRemoved)-1; i++ {
		if duplicateRemoved[i]+1 == duplicateRemoved[i+1] {
			temp++
			if temp > result {
				result = temp
			}
		} else {
			temp = 1
		}
	}
	return result
}
