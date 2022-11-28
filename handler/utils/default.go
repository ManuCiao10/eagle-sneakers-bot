package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/denisbrodbeck/machineid"
	"github.com/fatih/color"
)

func SelectMode(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func Gen_id() string {
	id, err := machineid.ProtectedID("myAppName")
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func Banner() {
	content, err := os.ReadFile("handler/utils/banner.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\033[H\033[2J")
	color.Red(string(content))
}

func Site_list() {
	color.Red("[Eagle 0.0.2] " + "[" + time.Now().Format("15:04:05.000000") + "]" + " 1. NEW BALANCE")
	color.Red("[Eagle 0.0.2] " + "[" + time.Now().Format("15:04:05.000000") + "]" + " 2. DADSTOCK")
	color.Red("[Eagle 0.0.2] " + "[" + time.Now().Format("15:04:05.000000") + "]" + " 3. KITH")
	color.Red("[Eagle 0.0.2] " + "[" + time.Now().Format("15:04:05.000000") + "]" + " 4. SUGAR")
	color.Red("[Eagle 0.0.2] " + "[" + time.Now().Format("15:04:05.000000") + "]" + " 5. SUSI")
	color.Red("[Eagle 0.0.2] " + "[" + time.Now().Format("15:04:05.000000") + "]" + " 6. 18 MONTROSE")

	println("\n")
}

func ConsolePrint(msg string, inputColor string) {
	switch inputColor {
	case "red":
		color.Red("[Eagle 0.0.2]" + "[" + time.Now().Format("15:04:05.000000") + "] " + msg)
	case "green":
		color.Green("[Eagle 0.0.2]" + "[" + time.Now().Format("15:04:05.000000") + "] " + msg)
	case "yellow":
		color.Yellow("[Eagle 0.0.2]" + "[" + time.Now().Format("15:04:05.000000") + "] " + msg)
	case "blue":
		color.Blue("[Eagle 0.0.2]" + "[" + time.Now().Format("15:04:05.000000") + "] " + msg)
	case "magenta":
		color.HiMagenta("[Eagle 0.0.2]" + "[" + time.Now().Format("15:04:05.000000") + "] " + msg)
	case "cyan":
		color.Cyan("[Eagle 0.0.2]" + "[" + time.Now().Format("15:04:05.000000") + "] " + msg)
	case "white":
		color.White("[Eagle 0.0.2]" + "[" + time.Now().Format("15:04:05.000000") + "] " + msg)
	case "black":
		color.Black("[Eagle 0.0.2]" + "[" + time.Now().Format("15:04:05.000000") + "] " + msg)
	default:
		color.Red("[Eagle 0.0.2]" + "[" + time.Now().Format("15:04:05.000000") + "] " + msg)
	}
}
