package ws_quicktasking

import (
	"fmt"
	"log"
	ps "github.com/mitchellh/go-ps"
)

func Initialize() {
	//check if the process is already running
	//if it is, then exit
	//if it isn't, then continue
	processList, err := ps.Processes()
	if err != nil {
		log.Println("ps.Processes() Failed, are you using windows?")
		return
	}

	// map ages
	for x := range processList {
		var process ps.Process
		process = processList[x]
		log.Printf("%d\t%s\n", process.Pid(), process.Executable())

		// do os.* stuff on the pid
	}
	fmt.Println("Connecting to quicktask websocket...")
	success := make(chan bool)
	go handleWebsocket(success)
	didSucceed := <-success
	if didSucceed {
		fmt.Println("Successfully authenticated to quicktask websocket.")
		//update Windows Title
	}
}
