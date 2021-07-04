package utils

import (
	"strconv"
)

func IsInt(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}
	return true
}
