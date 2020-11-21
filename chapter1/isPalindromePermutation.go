package chapter1

func isPalindromePermutation(s string) bool {

	charMap := make(map[rune]int)

	for _, char := range s {

		if char != ' ' {
			charMap[char]++
		}
	}

	var foundOdd bool
	for _, count := range charMap {
		if count%2 == 1 {
			if foundOdd {
				return false
			}

			foundOdd = true
		}

	}

	return true

}
