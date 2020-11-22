package chapter1

import (
	"fmt"
)

// abcde abcd
func IsOneEditDistance(check string, compare string) bool {

	var inPlaceIndex int
	var upIndex int
	var downIndex int

	var inPlaceMissCount int
	var upMissCount int
	var downMissCount int

	//Ensure that strings are in order now
	if len(check) > len(compare) {
		temp := check
		check = compare
		compare = temp
	}

	for _, char := range compare {
		if inPlaceIndex > len(check)-1 || char != []rune(check)[inPlaceIndex] {
			inPlaceMissCount++
		}

		if upIndex > len(check)-1 || char != []rune(check)[upIndex] {
			upMissCount++
			upIndex++
		}
		if downIndex > len(check)-1 || char != []rune(check)[downIndex] {
			downMissCount++
			downIndex--
		}

		inPlaceIndex++
		upIndex++
		downIndex++
	}

	fmt.Println(inPlaceMissCount)
	fmt.Println(upMissCount)
	fmt.Println(downMissCount)
	if inPlaceMissCount > 1 && upMissCount > 1 && downMissCount > 1 {
		return false
	}

	return true
}
