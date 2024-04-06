package main

import (
	"fmt"
	
)

func main() {
  s := []int{1,2,3,4,5,6,7,8}
  fmt.Print(twoSum(s,11))
}
func twoSum(numbers []int, target int) []int {
  i, j := 0, len(numbers) - 1
  sum := numbers[i] + numbers[j]
  
  fmt.Print(sum)
    
  fmt.Print("\n")
  for sum != target {

      if sum < target {
          i++
      } else {
          j--
      }
      
      sum = numbers[i] + numbers[j]
      fmt.Print(sum)
      
  }
  
  return []int{i + 1, j + 1}
}