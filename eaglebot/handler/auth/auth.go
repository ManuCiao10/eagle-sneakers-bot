package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"os"

	// "net/http"
	"os/user"
	"strings"

	"github.com/jaypipes/ghw"
)

func newSHA256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func GenerateHWID() string {
	block, _ := ghw.Block()
	var disks []string
	for _, disk := range block.Disks {
		disks = append(disks, disk.SerialNumber)
	}

	userStruct, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	username := userStruct.Username

	return newSHA256(strings.Join(disks, ",") + "," + username)
}

func Initialize() {
	//check for updates
	//create a .exe file
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	http.HandleFunc("/", helloHandler)
	log.Println("Listing for " + port)
	log.Print(http.ListenAndServe(":"+port, nil))
}
