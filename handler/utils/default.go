package utils

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/eagle/handler/version"
	"github.com/fatih/color"
)

func Menu() {
	mode := SelectMode(color.MagentaString("[Eagle " + version.Version + "]" + " [" + time.Now().Format("15:04:05.000000") + "]" + color.WhiteString(" PLESE SELECT A SITE:")))
	if mode == "1" {
		print("GAMESTOP")
	} else if mode == "2" {
		print("DEADSTOCK")
	} else if mode == "3" {
		print("TEST ZARA")
	} else {
		ConsolePrint("INVALID OPTION!", "red")
	}
}

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

//go:embed banner.txt
var content embed.FS

func Banner() {

	banner, _ := content.ReadFile("banner.txt")

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	color.Magenta(string(banner))
}

func Site() {
	version.Version = version.ExecutableName()
	color.Magenta("[Eagle " + version.Version + "] " + "[" + time.Now().Format("15:04:05.000000") + "]" + color.WhiteString(" 1. NEW BALANCE"))
	color.Magenta("[Eagle " + version.Version + "] " + "[" + time.Now().Format("15:04:05.000000") + "]" + color.WhiteString(" 2. DADSTOCK"))
	color.Magenta("[Eagle " + version.Version + "] " + "[" + time.Now().Format("15:04:05.000000") + "]" + color.WhiteString(" 3. KITH"))
	color.Magenta("[Eagle " + version.Version + "] " + "[" + time.Now().Format("15:04:05.000000") + "]" + color.WhiteString(" 4. SUGAR"))
	color.Magenta("[Eagle " + version.Version + "] " + "[" + time.Now().Format("15:04:05.000000") + "]" + color.WhiteString(" 5. SUSI"))
	color.Magenta("[Eagle " + version.Version + "] " + "[" + time.Now().Format("15:04:05.000000") + "]" + color.WhiteString(" 6. 18 MONTROSE"))

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
