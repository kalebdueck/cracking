package chapter1

import (
	"unicode"
)

func Urlify(s []rune) string {

	//Get a count of all spaces, that way we know how big to turn s into
	var stringCount int
	for _, runeVal := range s {
		if runeVal == ' ' {
			stringCount++
		}
	}

	oldLength := len(s)
	//there's a difference of 2 chars between a space and %20
	newLength := len(s) + stringCount*2

	//Cant just magically grow a slice in go
	//Gotta append some zero vals to the s slice
	sliceAddition := newLength - len(s)
	appendSlice := make([]rune, sliceAddition)

	s = append(s, appendSlice...)
	//gotta go backwards
	for i := oldLength - 1; i >= 0; i-- {
		if unicode.IsSpace(s[i]) {
			s[newLength-1] = '0'
			s[newLength-2] = '2'
			s[newLength-3] = '%'

			newLength -= 3
		} else {
			s[newLength-1] = s[i]
			newLength--
		}
	}

	return string(s)
}
