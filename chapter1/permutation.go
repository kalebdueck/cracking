package chapter1

func isPermutation(a string, b string) bool {

	aMap := mapStringCharCount(a)
	bMap := mapStringCharCount(b)

	for char, count := range aMap {
		if count != bMap[char] {
			return false
		}

	}

	return true
}

func mapStringCharCount(newString string) map[rune]int {

	charMap := make(map[rune]int)
	for _, char := range newString {
		charMap[char]++
	}

	return charMap
}
