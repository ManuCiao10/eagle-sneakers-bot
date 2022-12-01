package console

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/eagle/handler/version"
)

var (
	carts     int
	checkouts int
	failures  int
)

func Initialize() {
	carts = 0
	checkouts = 0
	failures = 0

	updateTitle()
}

func SetConsoleTitle(title string) (int, error) {
	handle, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		return 0, err
	}
	defer syscall.FreeLibrary(handle)
	proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
	if err != nil {
		return 0, err
	}
	r, _, err := syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
	return int(r), err
}

func updateTitle() {
	SetConsoleTitle(fmt.Sprintf("Eagle - EagleBot Version %s ", version.Version))
	// _, _ = SetConsoleTitle(fmt.Sprintf("Eagle - EagleBot Version %d ｜ Carts: %d ｜ Checkouts: %d ｜ Failures: %d",version.Version, carts, checkouts, failures))
}
