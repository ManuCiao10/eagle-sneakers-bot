package console

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"

	"github.com/eagle/handler/version"
)

var (
	carts     int
	checkouts int
	failures  int
)

func AddCheckout() {
	checkouts += 1
	updateTitle()
}

func AddCart() {
	carts += 1
	updateTitle()
}

func AddFailure() {
	failures += 1
	updateTitle()
}

func Initialize() {
	carts = 0
	checkouts = 0
	failures = 0

	updateTitle()
}

// setConsoleTitle
// func setConsoleTitle(title string) (int, error) {
// 	kernel32 := syscall.NewLazyDLL("kernel32.dll")
// 	proc := kernel32.NewProc("SetConsoleTitleW")
// 	ret, _, err := proc.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))))
// 	return int(ret), err

// }

// func updateTitle() {
// 	_, _ = setConsoleTitle(fmt.Sprintf("HellasAIO ｜ Carts: %d ｜ Checkouts: %d ｜ Failures: %d", carts, checkouts, failures))
// }