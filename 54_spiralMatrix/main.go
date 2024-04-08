package main

import "fmt"

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	// aa := minWindow(s, t)
	// fmt.Print(aa)
	hashMap := make(map[uint8]int)
	hashMap[t[0]]++
	fmt.Print(minWindow(s,t))


}


func spiralOrder(matrix [][]int) []int {
    rowTop :=1
    colLeft:=0
    rowNumber:=len(matrix)-1
    colNumber:=len(matrix[0])-1
    output:=[]int{}
    totalSpace:= rowNumber*colNumber


    col:=0
    row:=0
    direction:=true

    for i:=1;i<=totalSpace;i++{
        output = append(output,matrix[row][col])
        if(direction){
            if(col < colNumber){
                col++
            }else{
                row++
            }
            if(row>rowNumber){
                row--
                col--
                rowNumber--
                colNumber--
                direction=false
                continue
            }
        } else{

             if(col > colLeft){
                col--
            }else{
                row--
            }
            if(row<rowTop){
                row++
                col++
                rowTop++
                colLeft++
                direction=true
                continue
            }

        }


    }

    return output
}