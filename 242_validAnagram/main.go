package main

import (
	"fmt"
	"strings"
)

func main() {

	// aa := minWindow(s, t)
	// fmt.Print(aa)
	pattern := "abba"
	s := "dog cat cat dog"
	wordPattern(pattern, s)

}
func isAnagram(s string, t string) bool {
	count:=len(s)
	if count !=len(t){
		return false
	}
	hashMap := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		_, hasKey := hashMap[s[i]]
		if hasKey {
			hashMap[s[i]]++
		} else {
			hashMap[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		_, hasKey := hashMap[t[i]]
		if !hasKey{
			return false
		}else{
			hashMap[t[i]]--
			if	hashMap[t[i]]<0{
				return false
			}
		}
	


}
return true
}