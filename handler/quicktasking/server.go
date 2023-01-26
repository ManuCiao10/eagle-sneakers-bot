package quicktasking

import (
	"log"
	"net/http"

	"github.com/shirou/gopsutil/process"
)

var DevMode = true

func start() {
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
			// fmt.Println("Eagle is already running. (", name, ")")
			// time.Sleep(2 * time.Second)
			return
		}
	}

	// console.AddMQT()
	http.HandleFunc("/quicktask", quicktaskHandler)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
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
