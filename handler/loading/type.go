package loading

import (
	"bufio"
	"strings"
)

func CreateSliceProxy(scanner *bufio.Scanner) []string {
	var proxies []string

	for scanner.Scan() {
		proxies = append(proxies, scanner.Text())
	}

	return proxies
}

func Trim(s string) string {
	return strings.TrimSpace(s)
}
