package util

import "strings"

func ContainsIgnoreCase(slice []string, val string) bool {
	if len(slice) == 0 {
		return false
	}

	for _, value := range slice {
		if strings.EqualFold(value, val) {
			return true
		}
	}

	return false
}
