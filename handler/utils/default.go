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

	"github.com/eagle/handler/logs"
	"github.com/eagle/handler/version"
	"github.com/fatih/color"
)

var (
	Debug        = false
	Dev          = true
	THEBROKENARM = 1
	MONITOR      = 2
	Active       bool

	ERROR = 255
)

// add grid sites
func Menu() int {
	mode := SelectMode(color.MagentaString(version.GetVersion() + logs.Time() + color.WhiteString("PLESE SELECT A SITE:")))
	if mode == "1" {
		return THEBROKENARM
	} else if mode == "2" {
		return MONITOR
	}
	return ERROR
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
	color.Magenta(version.GetVersion() + logs.Time() + color.WhiteString("1. THEBROKENARM"))
	color.Magenta(version.GetVersion() + logs.Time() + color.WhiteString("2. EAGLE MONITOR"))

	println("\n")
}

func Directory(site string) {
	var i = 1
	files, err := os.ReadDir("./" + site)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.Name() != "accounts.csv" {
			s := strconv.Itoa(i)
			color.Magenta(version.GetVersion() + logs.Time() + color.WhiteString(s+". "+f.Name()))
			i++
		}

	}
	println("\n")
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if strings.Contains(v, str) {
			return true
		}
	}
	return false
}

func ProxyToUrl(proxy string) string {
	proxySplit := strings.Split(proxy, ":")

	if len(proxySplit) == 2 {
		return fmt.Sprintf("http://%s:%s", proxySplit[0], proxySplit[1])
	} else if len(proxySplit) == 4 {
		return fmt.Sprintf("http://%s:%s@%s:%s", proxySplit[2], proxySplit[3], proxySplit[0], proxySplit[1])
	}

	return fmt.Sprintf("http://%s", proxy)
}
