package chapter1

import (
	"fmt"
	"strings"
)

type CharCount struct {
	char  rune
	count int
}

//So my initial naive thought is to hold a slice of strcuts of char && count
func CompressString(s string) string {
	var CharCountList []CharCount

	for i := 0; i < len(s); {
		CharCountList = append(CharCountList, CharCount{char: []rune(s)[i], count: 1})

		var j int
		for j = i + 1; j < len(s); j++ {
			if []rune(s)[j] != []rune(s)[i] {
				break
			}

			CharCountList[len(CharCountList)-1].count++
		}

		i = j
	}

	var sb strings.Builder
	for _, CharCounts := range CharCountList {
		sb.WriteRune(CharCounts.char)
		sb.WriteString(fmt.Sprint(CharCounts.count))
	}

	if len(sb.String()) > len(s) {
		return s
	}
	return sb.String()
}

//Don't need a map or slice or anything
func CompressStringBookMethod(s string) string {
	var sb strings.Builder
	var duplicateCount int
	for i, char := range s {
		duplicateCount++

		if i+1 >= len(s) || []rune(s)[i+1] != char {
			sb.WriteRune([]rune(s)[i])
			sb.WriteString(fmt.Sprint(duplicateCount))
			duplicateCount = 0
		}
	}

	if len(sb.String()) > len(s) {
		return s
	}
	return sb.String()
}
