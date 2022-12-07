package thebrokenarm

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func Write_data_to_file(data string, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		// Print_err("FILE CREATION ERROR")
	}
	defer f.Close()
	f.WriteString(data)
}

func RandomString(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
