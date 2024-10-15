package strutil

import "unicode"

// IsBlank checks if string is empty ("") or whitespace only.
func IsBlank(s string) bool {
	if s == "" {
		return true
	}

	// checks whitespace only
	for _, v := range s {
		if !unicode.IsSpace(v) {
			return false
		}
	}

	return true
}

// IsEmpty alias of len(s) == 0 || s == ""
func IsEmpty(s string) bool {
	return s == ""
}
