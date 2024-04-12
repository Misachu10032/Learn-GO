package main

import (
	"fmt"
	"sort"
)

func main() {

	// aa := minWindow(s, t)
	// fmt.Print(aa)
	strs := []int{1,2,3,4,5,6,7,8,9,9}
	fmt.Print(containsNearbyDuplicate(strs,3))

}
func containsNearbyDuplicate(nums []int, k int) bool {
    if(k>=len(nums)){
        array := nums
        sort.Ints(array)
        for i:=0;i<len(nums)-1;i++{
            if array[i]==array[i+1]{
                return true
            }
        }
        return false
    }
    
    for i:=0;i<len(nums)-k;i++{
        for j:=1;j<=k;j++{
			fmt.Print("\n")
			fmt.Print(nums[i],"   ",nums[i+j])
            if nums[i]==nums[i+k]{
                return true
            }
        }
    }
    return false
}