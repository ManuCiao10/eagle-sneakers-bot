package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/eagle/eaglebot/handler/loading"

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

//VERSION
// func Initialize() {
// 	//check for updates
// 	//create a .exe file
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "9000"
// 	}
//	mux := http.NewServeMux()
// 	helloHandler := func(w http.ResponseWriter, req *http.Request) {
// 		io.WriteString(w, "API is running\n")
// 	}
// 	DownloadHandler := func(w http.ResponseWriter, req *http.Request) {
// 		//download the file.exe
// 		return version.Download(w, req)
// 	}

// 	VersionHandler := func(w http.ResponseWriter, req *http.Request) {
// 		//return version
// 		println("Checking for updates...")
// 		resp, err := http.Get("https://eagleaio.herokuapp.com/version")
// 		if err != nil {
// 			println("Error checking for updates")
// 			return
// 		}
// 		defer resp.Body.Close()
// 		newversion := resp.Body
// 		if newversion == version.Version {
// 			println("You are up to date")
// 		} else {
// 			println("New version available")
// 			version.Update()
// 		}

// 		// io.WriteString(w, version.Version)

// 	}
// 	http.HandleFunc("/", helloHandler)
// 	http.HandleFunc("/version", VersionHandler)
// 	http.HandleFunc("/download", DownloadHandler)

// 	log.Print(http.ListenAndServe(":"+port, nil))
// }

// KEY
func Initialize() {
	if len(loading.Data.Settings.Settings.AuthKey) != 300 {
		println("[!] Invalid key")
		time.Sleep(2 * time.Second)
		os.Exit(1)
	}
}
