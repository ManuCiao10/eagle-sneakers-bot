package auth

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/eagle/eaglebot/handler/loading"
	"github.com/eagle/eaglebot/handler/utils"
	"github.com/fatih/color"

	"github.com/jaypipes/ghw"
)

var Auth AuthResponse

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

func ValidateHWID(key string) bool {
	HWID := GenerateHWID()
	check_id := Auth.Metadata.HWID //read from json struct

	if check_id == "" {
		url := "https://api.hyper.co/v6/licenses/" + key + "/metadata"
		payload, err := json.Marshal(map[string]interface{}{
			"metadata" : map[string]interface{}{
				"hwid" : HWID,
			},
		})

		if err != nil {
			log.Fatal(err)
		}

		req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(payload))

		req.Header.Add("content-type", "application/json")
		req.Header.Add("Authorization", "Bearer "+os.Getenv("API_TOKEN"))

		res, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		fmt.Println(res)
		fmt.Println(string(body))
	} else if check_id != HWID {
		color.Yellow("[" + time.Now().Format("15:04:05.000000") + "] " + "MACHINE ALREADY BOUND")
		time.Sleep(3 * time.Second)
		os.Exit(255)
	}
	color.Yellow("[" + time.Now().Format("15:04:05.000000") + "] " + "UNABLE TO VALIDATE MACHINE ID")
	return false

}

func ValidateKey(key string) bool {
	client := &http.Client{}
	color.Yellow("[" + time.Now().Format("15:04:05.000000") + "] " + "VALIDATING KEY...")

	r, err := http.NewRequest("GET", "https://api.hyper.co/v6/licenses/"+key, nil)
	if err != nil {
		utils.ConsolePrint("AUTH SERVER ERROR", "red")
		return false
	}

	r.Header.Set("Authorization", "Bearer "+os.Getenv("API_TOKEN"))
	r.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		utils.ConsolePrint("COULD NOT REQUEST AUTH", "red")
		return false
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.ConsolePrint("COULD NOT READ RESPONSE BODY", "red")
		return false
	}

	if resp.StatusCode != 200 {
		return false
	}

	err = json.Unmarshal(body, &Auth)
	if err != nil {
		log.Fatal("Could not unmarshal response body, shutting down.", err)
	}

	ValidateHWID(key)

	return true

}

// KEY
func Initialize() {
	//check for updates
	if len(loading.Data.Settings.Settings.AuthKey) == 0 {
		color.HiMagenta("[" + time.Now().Format("15:04:05.000000") + "] " + "INVALID KEY DETECTED [CHECK JSON FILE]")
		time.Sleep(2 * time.Second)
		os.Exit(1)
	}
	key := loading.Data.Settings.Settings.AuthKey

	if !ValidateKey(key) {
		color.HiRed("[" + time.Now().Format("15:04:05.000000") + "] " + "INVALID KEY")
		time.Sleep(2 * time.Second)
		os.Exit(1)
	}
}
