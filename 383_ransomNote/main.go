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
func canConstruct(ransomNote string, magazine string) bool {
	count := len(ransomNote)
	hashMap := make(map[uint8]int)
	for i := 0; i < count; i++ {

		_, hasKey := hashMap[ransomNote[i]]
		if hasKey {
			hashMap[ransomNote[i]]++
		} else {
			hashMap[ransomNote[i]] = 1
		}
	}

	for i := 0; i < len(magazine); i++ {
		value, hasKey := hashMap[magazine[i]]
		if hasKey && value > 0 {
			hashMap[magazine[i]]--
			count--
			if count == 0 {
				return true
			}
		}

	}

	return false
}
