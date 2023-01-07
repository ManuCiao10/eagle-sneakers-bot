package task

import (
	"math/rand"
	"strings"
)

func RandomSize() string {
	sizes := []string{"36", "36.5", "37", "38"}

	return sizes[rand.Intn(len(sizes))]
}

func SplitSize(size string) string {
	sizes := strings.Split(size, ";")

	return sizes[rand.Intn(len(sizes))]
}

func Contains(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}
