package main
import (
	"fmt"
    "sort"
)
func threeSum(nums []int) [][]int {
    sortNums := nums
    sort.Ints(sortNums)
    output := make([][]int,0)
    i,j:=0,len(nums)-1
    for i<j{
        sum:=sortNums[i]+sortNums[j]
        for k:=i+1;k<j-1;k++{


            if(sum+sortNums[k]==0){
                 output = append(output, []int{sortNums[i], sortNums[j], sortNums[k]})
                
            }

        }
        i++
        j--
    }
    return output

}


func sum() {
	s := []int{-1,0,1,2,-1,-4}
fmt.Print(threeSum(s))
  }