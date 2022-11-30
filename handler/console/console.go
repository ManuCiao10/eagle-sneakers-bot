package console

import (
	"time"

	"github.com/eagle/handler/version"
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

// func setConsoleTitle(title string) (int, error) {
// 	handle, err := syscall.LoadLibrary("Kernel32.dll")
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer syscall.FreeLibrary(handle)
// 	proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
// 	if err != nil {
// 		return 0, err
// 	}
// 	r, _, err := syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
// 	return int(r), err
// }

// func updateTitle() {
// 	_, _ = setConsoleTitle(fmt.Sprintf("HellasAIO ｜ Carts: %d ｜ Checkouts: %d ｜ Failures: %d", carts, checkouts, failures))
// }
