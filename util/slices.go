package util

import "strings"

func ContainsIgnoreCaseAndWithoutSpace(slice []string, val string) bool {
	if len(slice) == 0 {
		return false
	}

	for _, value := range slice {
		if EqualFoldWithoutSpace(value, val) {
			return true
		}
	}

	return false
}

func EqualFoldWithoutSpace(s1 string, s2 string) bool {
	return strings.EqualFold(strings.Replace(s1, " ", "", -1),
		strings.Replace(s2, " ", "", -1))
}
