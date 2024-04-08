package main

import "fmt"

func main() {
    s :="wordgoodgoodgoodbestword"
    words:=[]string{"word","good","best","good"}
    aa:=findSubstring(s,words)
    fmt.Print(aa)
}



func findSubstring(s string, words []string) []int {
	numOfWord := len(words)
	lenOfWord := len(words[0])
	requiredLen := numOfWord * lenOfWord
	hashMap := make(map[string]int)
	tempMap := make(map[string]int)

    result := make([] int,0)
	
	for i := 0; i < numOfWord; i++ {
		_, hasKey := hashMap[words[i]]
		if hasKey {
			hashMap[words[i]]++
		} else {
			hashMap[words[i]] = 1
		}
	}
	for i := 0; i < len(s)-requiredLen+1; i++ {

		for k := 0; k < numOfWord; k++ {
        tempWords:=""
			for j := 0; j < lenOfWord; j++ {
           
				tempWords += string(s[i+j+k*lenOfWord])
			}
			occurrence, hasKey := hashMap[tempWords]

            fmt.Print("\n")
            fmt.Print(tempWords)
           

			if hasKey {
				tempMap[tempWords]++
				if tempMap[tempWords] > occurrence {
					tempMap = make(map[string]int)
					break
				}
			}else {
                tempMap = make(map[string]int)
                break
            }

            if k==numOfWord-1{
                result = append(result,i)
            }

		}

	}

    return result
}