package main

import (
	"fmt"
)

func main() {

	// aa := minWindow(s, t)
	// fmt.Print(aa)
	strs := 7
	fmt.Print(isHappy(strs))

}
func isHappy(n int) bool {
	sumSquare := n
	hashMap := make(map[int]int)
	for sumSquare != 1 {
		temp := sumSquare
		tempSum := 0
		for temp >= 10 {
			modulus := temp % 10
			tempSum += modulus * modulus
			temp /= 10
		}
		fmt.Print(tempSum)
		sumSquare = tempSum
		_, ok := hashMap[sumSquare]
		if ok {
			return false
		} else {
			hashMap[sumSquare] = sumSquare
		}
	}
	return true

}
