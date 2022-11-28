package console

import (
	"time"

	"github.com/eagle/eaglebot/handler/version"
)

var (
	carts     int
	checkouts int
	failures  int
)

// func setConsoleTitle(title string) (int, error) {
// 	//set Console cli
// 	// handle := syscall.LoadLibrary
// 	return 0, fmt.Errorf(error)

// }

func Initialize() {
	carts = 0
	checkouts = 0
	failures = 0

	//For *NIX systems, print \033]0;Title goes here\007 to stdout
	print("\033]0;EagleBot Version " + version.Version + " ｜ Carts: " + string(rune(carts)) + " ｜ Checkouts: " + string(checkouts) + " ｜ Failures: " + string(failures) + "\007")
	//for Windows, use syscall.LoadLibrary and syscall.GetProcAddress to call SetConsoleTitleA
	// _, _ = setConsoleTitle(fmt.Sprintf("EagleBot Version %d ｜ Carts: %d ｜ Checkouts: %d ｜ Failures: %d", version.Version, carts, checkouts, failures))
	time.Sleep(5 * time.Second)

}
