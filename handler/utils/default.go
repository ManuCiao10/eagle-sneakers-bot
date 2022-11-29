package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eagle/handler/version"
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

func Banner() {
	content, err := os.ReadFile("handler/utils/banner.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\033[H\033[2J")
	color.Red(string(content))
}

func GetVersionName() {
	file, err := os.Open("EagleBot/")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()
	list, _ := file.Readdirnames(0)

	for _, name := range list {
		if strings.Contains(name, ".exe") {
			version_ := strings.Split(name, "_")[1]
			version.Version = version_[:len(version_)-4]
		}
	}
}

func Site_list() {
	color.Red("[Eagle " + version.Version + "] " + "[" + time.Now().Format("15:04:05.000000") + "]" + " 1. NEW BALANCE")
	color.Red("[Eagle " + version.Version + "] " + "[" + time.Now().Format("15:04:05.000000") + "]" + " 2. DADSTOCK")
	color.Red("[Eagle " + version.Version + "] " + "[" + time.Now().Format("15:04:05.000000") + "]" + " 3. KITH")
	color.Red("[Eagle " + version.Version + "] " + "[" + time.Now().Format("15:04:05.000000") + "]" + " 4. SUGAR")
	color.Red("[Eagle " + version.Version + "] " + "[" + time.Now().Format("15:04:05.000000") + "]" + " 5. SUSI")
	color.Red("[Eagle " + version.Version + "] " + "[" + time.Now().Format("15:04:05.000000") + "]" + " 6. 18 MONTROSE")

	println("\n")
}

func ConsolePrint(msg string, inputColor string) {
	switch inputColor {
	case "red":
		color.Red("[Eagle " + version.Version + "]" + " [" + time.Now().Format("15:04:05.000000") + "] " + msg)
	case "green":
		color.Green("[Eagle " + version.Version + "]" + " [" + time.Now().Format("15:04:05.000000") + "] " + msg)
	case "yellow":
		color.Yellow("[Eagle " + version.Version + "]" + " [" + time.Now().Format("15:04:05.000000") + "] " + msg)
	case "blue":
		color.Blue("[Eagle " + version.Version + "]" + " [" + time.Now().Format("15:04:05.000000") + "] " + msg)
	case "magenta":
		color.Magenta("[Eagle " + version.Version + "]" + " [" + time.Now().Format("15:04:05.000000") + "] " + msg)
	case "cyan":
		color.Cyan("[Eagle " + version.Version + "]" + " [" + time.Now().Format("15:04:05.000000") + "] " + msg)
	case "white":
		color.White("[Eagle " + version.Version + "]" + " [" + time.Now().Format("15:04:05.000000") + "] " + msg)
	case "black":
		color.Black("[Eagle " + version.Version + "]" + " [" + time.Now().Format("15:04:05.000000") + "] " + msg)
	default:
		color.Red("[Eagle " + version.Version + "]" + " [" + time.Now().Format("15:04:05.000000") + "] " + msg)
	}
}
