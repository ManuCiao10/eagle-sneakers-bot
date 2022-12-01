package auth

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/utils"

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
	check_id := Auth.Metadata.HWID

	if check_id == "" {
		url := "https://api.hyper.co/v6/licenses/" + key + "/metadata"
		payload, err := json.Marshal(map[string]interface{}{
			"metadata": map[string]interface{}{
				"hwid": HWID,
			},
		})

		if err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(payload))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Add("content-type", "application/json")
		req.Header.Add("Authorization", "Bearer "+loading.Data.Env.Env.API_TOKEN)

		_, err = http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}

	} else if check_id != HWID {
		color.White("[" + time.Now().Format("15:04:05.000000") + "] " + "RESET YOUR KEY IN THE DASHBOARD")
		time.Sleep(3 * time.Second)
		os.Exit(0)
	}

	return true

}

func ValidateKey(key string) bool {
	client := &http.Client{}

	color.Magenta("[" + time.Now().Format("15:04:05.000000") + "] " + "VALIDATING KEY...")

	r, err := http.NewRequest("GET", "https://api.hyper.co/v6/licenses/"+key, nil)
	if err != nil {
		utils.ConsolePrint("AUTH SERVER ERROR", "red")
		return false
	}
	time.Sleep(2 * time.Second)
	r.Header.Set("Authorization", "Bearer "+loading.Data.Env.Env.API_TOKEN)
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

	if !ValidateHWID(key) {
		return false
	}
	return true

}

func Initialize() {
	if len(loading.Data.Settings.Settings.AuthKey) == 0 {
		color.HiMagenta("[" + time.Now().Format("15:04:05.000000") + "] " + "INVALID KEY DETECTED [CHECK JSON FILE]")
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}
	key := loading.Data.Settings.Settings.AuthKey

	if !ValidateKey(key) {
		color.HiRed("[" + time.Now().Format("15:04:05.000000") + "] " + "INVALID KEY")
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}
}
