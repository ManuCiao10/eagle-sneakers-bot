package deadstock

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
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
		Print_err("FILE CREATION ERROR")
	}
	defer f.Close()
	f.WriteString(data)
}

func Print_err(msg string) {
	color.Red("[Eagle 0.0.2]"+"["+time.Now().Format("15:04:05.000000")+"]"+" %s", msg)
	os.Exit(0)
}

func Print(msg string) {
	color.Magenta("[Eagle 0.0.2]"+"["+time.Now().Format("15:04:05.000000")+"]"+" %s", msg)
}

func Print_cart(msg string) {
	color.Cyan("[Eagle 0.0.2]"+"["+time.Now().Format("15:04:05.000000")+"]"+" %s", msg)
}

func Print_err_cart(msg string) {
	color.Yellow("[Eagle 0.0.2]"+"["+time.Now().Format("15:04:05.000000")+"]"+" %s", msg)
}

func randomString(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}