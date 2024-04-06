func lengthOfLongestSubstring(s string) int {
	hashMap := make(map[uint8]int)
	length := 0
	startingIndex := 0
	output := 0
	for i := 0; i < len(s); i++ {
		_, hasKey := hashMap[s[i]]
		if hasKey {
			if s[startingIndex] == s[i] {
				startingIndex++
				hashMap[s[i]] = i
			} else {
				startingIndex = hashMap[s[i]]
				hashMap = make(map[uint8]int)
				length = 1
				for j := startingIndex + 1; j <= i; j++ {
					hashMap[s[j]] = j
				}

			}

		} else {
			hashMap[s[i]] = i
			length++
			if length > output {
				output = length
			}
		}

	}

	return output
}