package main

import (
	"fmt"
	"strings"
)

func main() {

	// aa := minWindow(s, t)
	// fmt.Print(aa)
	pattern :="abba"
	s :="dog cat cat dog"
	wordPattern(pattern,s)

}
func wordPattern(pattern string, s string) bool {
    
	w := strings.Split(s, " ")
	sMap := make(map[uint8]string)
	wMap := make(map[string]uint8)

	if len(w)!=len(pattern){
		return false
	}
	for i:=0;i<len(w);i++{
		sValue,hasSKey:= sMap[pattern[i]]
		wValue,hasWKey:= wMap[w[i]]
		if hasSKey{
			if sValue != w[i]{
				fmt.Print(sValue,w[i],"ds")
				return false
			}
		}else{
			sMap[pattern[i]]=w[i]
		   wMap[w[i]]=pattern[i]
		}
		if hasWKey{
			fmt.Print(wValue,pattern[i],"ss")
			  if wValue != pattern[i]{
				return false
			}
		}
		fmt.Print("\n")

		fmt.Print(sMap,wMap,"map")

	}
	return true
}