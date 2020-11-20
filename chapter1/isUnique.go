package chapter1

func isUnique(s string) bool {
	charMap := make(map[rune]bool)

	for _, char := range s {
		_, exists := charMap[char]
		if exists == true {
			return false
		}

		charMap[char] = true
	}

	return true
}

func isUniqueSavingSpace(s string) bool {
	for i, char := range s {
		//no need to look at chars we've already looked at
		for j := i + 1; j < len(s); j++ {
			if char == []rune(s)[j] {
				return false
			}
		}
	}

	return true
}
