package chapter1

func IsStringRotation(s1 string, s2 string) bool {
	if len(s2) != len(s1) {
		return false
	}

	//If you mash a rotated string together, the initial string will be in there.
	//Cheeky.
	return IsSubstring(s1+s1, s2)
}
