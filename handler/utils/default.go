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

	// "github.com/eagle/handler/version"
	"github.com/eagle/handler/version"
	"github.com/fatih/color"
)

var (
	Debug        = false
	Dev          = true
	THEBROKENARM = 1
	MONITOR      = 2

	ERROR     = 255
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"
)

func Menu() int {
	mode := SelectMode(color.MagentaString(Version() + Time() + color.WhiteString("PLESE SELECT A SITE:")))
	if mode == "1" {
		return THEBROKENARM
	} else if mode == "2" {
		return MONITOR
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
	ConsolePrint(color.WhiteString("1. THEBROKENARM"), "magenta")
	ConsolePrint(color.WhiteString("2. EAGLE MONITOR"), "magenta")

	println("\n")
}

func ConsolePrint(msg string, inputColor string) {
	switch inputColor {
	case "red":
		color.Red(Version() + Time() + msg)
	case "green":
		color.Green(Version() + Time() + msg)
	case "yellow":
		color.Yellow(Version() + Time() + msg)
	case "blue":
		color.Blue(Version() + Time() + msg)
	case "magenta":
		color.Magenta(Version() + Time() + msg)
	case "cyan":
		color.Cyan(Version() + Time() + msg)
	case "white":
		color.White(Version() + Time() + msg)
	case "black":
		color.Black(Version() + Time() + msg)
	default:
		color.Red(Version() + Time() + msg)
	}
}

func Time() string {
	return "[" + time.Now().Format("15:04:05.000000") + "] "
}

func Version() string {
	return "[Eagle " + version.Version + "] "
}

func Directory(site string) {
	files, err := os.ReadDir("./" + site)
	if err != nil {
		log.Fatal(err)
	}
	for i, f := range files {
		i = i + 1
		s := strconv.Itoa(i)
		ConsolePrint(color.WhiteString(s+". "+f.Name()), "magenta")
	}
	println("\n")
}
