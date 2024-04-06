package main
import (
	"fmt"
    "sort"
)
func threeSum(nums []int) [][]int {
    sortNums := nums
    sort.Ints(sortNums)
    output := make([][]int,0)
    fmt.Print(sortNums)

    i,j:=0,len(nums)-1

    for i<j{
        sum:=sortNums[i]+sortNums[j]
      
        for k:=i+1;k<j;k++{
            fmt.Print("\n")
            fmt.Print(sum,"sum")
            fmt.Print( sortNums[k],"k")
           
            if(sum+sortNums[k]==0){
                 output = append(output, []int{sortNums[i], sortNums[j], sortNums[k]})
                
            }

        }
        i++
        j--
    }

    return output

}


func main() {
	s := []int{-1,0,1,2,-1,-4}
fmt.Print(threeSum(s))
  }