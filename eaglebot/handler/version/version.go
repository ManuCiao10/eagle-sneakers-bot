package version

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	Version = "0.0.17"
	File    Update
)


func CheckSum() string {
	sha256_hash := sha256.New()

	if _, err := os.Stat("bin/bot.exe"); os.IsNotExist(err) {
		if _, err := os.Stat("bin"); os.IsNotExist(err) {
			os.Mkdir("bin", 0755)
		}
		return ""
	}

	file, err := os.Open("bin/bot.exe")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if _, err := io.Copy(sha256_hash, file); err != nil {
		log.Fatal(err)
	}

	return string(sha256_hash.Sum(nil))
}

func GetLatestVersion() string {
	url := "https://api.hyper.co/v6/products"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("API_TOKEN"))

	res, _ := http.DefaultClient.Do(req)

	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	err := json.Unmarshal(body, &File)
	if err != nil {
		log.Fatal("Could not unmarshal response body, shutting down.", err)
	}
	for _, v := range File.Info {
		if strings.Contains(v.Version, "Update") {

			return v.Version
		}
	}
	return Version
}

func GetID() string {
	for _, v := range File.Info {
		if strings.Contains(v.Version, "Update") {
			return v.ID
		}
	}
	return ""
}

func DowloadUpdate() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://manuciao5388.hyper.co/ajax/products/"+ GetID() + "/files?", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("cookie", "authorization="+os.Getenv("AUTH_DOWNLOAD"))
	req.Header.Set("hyper-account", os.Getenv("ACC_DOWLOAD"))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat("bin"); os.IsNotExist(err) {
		os.Mkdir("bin", 0755)
	}
	final_version := GetLatestVersion()
	file, err := os.Create("bin/EagleBot_" + strings.Split(final_version, " ")[1] + ".exe")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.Write(bodyText)
	if err != nil {
		log.Fatal(err)
	}
	file.Sync()
	fmt.Print(resp.StatusCode)
}

func Updates() {
	new_version := strings.ToUpper(GetLatestVersion())
	new := strings.Split(new_version, ".")
	if new[2] != CheckSum() {
		fmt.Print(CheckSum())
		time.Sleep(30 * time.Second)
		color.White("[" + time.Now().Format("15:04:05.000000") + "] " + "DOWNLOADING " + new_version)
		DowloadUpdate()
		os.Exit(255)

	}
	time.Sleep(30 * time.Second)

}

//fix if wher you check the name of the currently bot