package ws_quicktasking

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/v3/process"
)

var DevMode = true

func Initialize() {
	executableName := "Eagle"
	if DevMode {
		executableName = "eagle"
	}

	processes, err := process.Processes()
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, process := range processes {
		name, _ := process.Name()
		if Contains(name, executableName) {
			return
		}
	}

	fmt.Println("Connecting to quicktask...")
	success := make(chan bool)
	go handleWebsocket(success)
	didSucceed := <-success
	if didSucceed {
		fmt.Println("Successfully authenticated to quicktask websocket.")
	}

}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if len(s)-i < len(substr) {
			return false
		}
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
