package utils

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/eagle/handler/version"
	"github.com/fatih/color"
)

var (
	THEBROKENARM = 1
	ERROR        = 255
)

func Menu() int {
	mode := SelectMode(color.MagentaString("[Eagle " + version.Version + "]" + " [" + time.Now().Format("15:04:05.000000") + "]" + color.WhiteString(" PLESE SELECT A SITE:")))
	if mode == "1" {
		return THEBROKENARM
	}
	return 255
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
	color.Magenta("[Eagle " + version.Version + "] " + "[" + time.Now().Format("15:04:05.000000") + "]" + color.WhiteString(" 1. THEBROKENARM"))

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

func Directory(site string) {
	files, err := os.ReadDir("./" + site)
	if err != nil {
		log.Fatal(err)
	}
	for i, f := range files {
		i = i + 1
		s := strconv.Itoa(i)
		color.Magenta("[Eagle " + version.Version + "]" + " [" + time.Now().Format("15:04:05.000000") + "] " + color.WhiteString(s+". "+f.Name()))
	}
	println("\n")
}
