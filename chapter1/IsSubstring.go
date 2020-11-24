package chapter1

func IsSubstring(s1 string, s2 string) bool {
	//Edge case
	if len(s2) > len(s1) {
		return false
	}

	for initial, _ := range s1 {
		for index, compare := range s2 {
			if []rune(s1)[initial+index] != compare {
				break
			}

			if index == len(s2)-1 {
				return true
			}
		}

	}
	return false
}
