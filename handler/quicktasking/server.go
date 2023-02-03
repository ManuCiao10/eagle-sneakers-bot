package quicktasking

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/shirou/gopsutil/process"
)

var DevMode = true
var count = 0

func start() {
	path := os.Args[0]
	executableName := filepath.Base(path)

	processes, err := process.Processes()
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, process := range processes {
		name, _ := process.Name()
		if Contains(name, executableName) {
			count++
		}
	}
	if count > 1 {
		return
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
