package stringsplus

import "strings"

// EqualFoldHasSuffix returns true if s has suffix, not considering case sensitivity.
func EqualFoldHasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && strings.EqualFold(s[len(s)-len(suffix):], suffix)
}

// EqualFoldHasPrefix returns true if s has prefix, not considering case sensitivity.
func EqualFoldHasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && strings.EqualFold(s[:len(prefix)], prefix)
}

// EqualFoldStringInSlice returns true if needle is in haystack, not considering case sensitivity.
func EqualFoldStringInSlice(needle string, haystack []string) bool {
	for _, hay := range haystack {
		if strings.EqualFold(hay, needle) {
			return true
		}
	}
	return false
}

// StringInSlice returns true if needle is in haystack.
func StringInSlice(needle string, haystack []string) bool {
	for _, hay := range haystack {
		if hay == needle {
			return true
		}
	}
	return false
}
