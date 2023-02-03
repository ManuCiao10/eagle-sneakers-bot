package utils

import (
	"regexp"
	"strings"
)

func GetId(body string) string {
	regex := regexp.MustCompile(`r:'(.*)'`)
	orderId := regex.FindStringSubmatch(body)
	if len(orderId[1]) > 1 {
		orderId := strings.Split(orderId[1], "'")
		return orderId[0]
	}

	return ""
}
