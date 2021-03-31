package numutil

func IsNum(str string) bool {
	for _, c := range []byte(str) {
		if !isNumCharacter(c) {
			return false
		}
	}
	return true
}

func isNumCharacter(c byte) bool {
	return '0' <= c && c <= '9'
}
