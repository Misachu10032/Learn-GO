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


func minWindow(s string, t string) string {

	if len(s) < len(t) {
		return ""
	}
	hashMap := make(map[uint8]int)
	tempMap := make(map[uint8]int)

	for i := 0; i < len(t); i++ {
		hashMap[t[i]]++
	}

	fmt.Print(hashMap)

	count := len(t)
	start, end := 0, 0
	windowLen := len(s) + 1
	startIndex := 0
	for end < len(s) {
		occurence,hasKey := hashMap[s[end]]
		if hasKey  {
			tempMap[s[end]]++
			if tempMap[s[end]] <= occurence {
				count--
			}

		}
		end++
		for count == 0 {
			if end-start < windowLen {

				windowLen = end - start
				startIndex=start
			}
			occurence,hasKey := hashMap[s[start]]
			if hasKey {
				tempMap[s[start]]--
				if tempMap[s[start]] < occurence {
					count++
				}

			}
            start++

		}

	}
	if windowLen == len(s)+1 {
		return ""
	}

	return s[startIndex : startIndex+windowLen]
}
