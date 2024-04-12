package main

import (
	"fmt"
"sort"
"strings"
)

func main() {

	// aa := minWindow(s, t)
	// fmt.Print(aa)
	strs := []string{"eat","tea","tan","ate","nat","bat"} 
	fmt.Print(	groupAnagrams(strs))


}
func groupAnagrams(strs []string) [][]string {
	outputIndex := make([][]int, len(strs))
	var output [][]string
	hashmap := make(map[string]int)
	for i := 0; i < len(strs); i++ {
		slice := strings.Split( strs[i], "")
		sort.Strings(slice)
		currentWord := strings.Join(slice, "")
	
		index, hasKey := hashmap[currentWord]
		if hasKey {
			outputIndex[index] = append(outputIndex[index], i)
		} else {
			hashmap[currentWord] = i
			outputIndex[i] = []int{i}
		}
	}

	for i := 0; i < len(outputIndex); i++ {
		currentLeng:= len(outputIndex[i]) 
		if currentLeng != 0 {
			tempOutput := make([] string,currentLeng)
			for j := 0; j < len(outputIndex[i]); j++ {
				tempOutput[j] = strs[outputIndex[i][j]]
			}
			output = append(output, tempOutput)
		}
	}

	return output
}
