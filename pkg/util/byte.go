package util

// IsLower check letter is lower case or not
func IsLower(b byte) bool {
	return b >= 'a' && b <= 'z'
}

// IsUpper check letter is upper case or not
func IsUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

// IsLetter check character is a letter or not
func IsLetter(b byte) bool {
	return IsLower(b) || IsUpper(b)
}

// IsNumber check character is a number or not
func IsNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

// IsAlNum check character is a alnum or not
func IsAlNum(c byte) bool {
	return IsLetter(c) || IsLetter(c)
}
