package version

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	Version string
	File    Update
)

func CheckSum() string {
	//no Folder bin => create
	if _, err := os.Stat("bin"); os.IsNotExist(err) {
		os.Mkdir("bin", 0755)
		return ""
	}

	file, err := os.Open("bin/")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()
	var count int
	list, _ := file.Readdirnames(0)

	//More files.exe => delete
	for _, name := range list {
		if strings.Contains(name, ".exe") {
			count++
		}
	}
	if count > 1 {
		color.Red("[" + time.Now().Format("15:04:05.000000") + "] " + "DELETE OLD VERSION")
		time.Sleep(2 * time.Second)
		os.Exit(255)
	}

	for _, name := range list {
		if strings.Contains(name, ".exe") {
			version := strings.Split(name, "_")[1]
			//save version in struct
			return version[:len(version)-4]
		}
	}
	return ""
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
	return ""
}

func GetID() string {
	for _, v := range File.Info {
		if strings.Contains(v.Version, "Update") {
			return v.ID
		}
	}
	return ""
}

func DowloadUpdate() bool {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://manuciao5388.hyper.co/ajax/products/"+GetID()+"/files?", nil)
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

	return resp.StatusCode == 200
}

func Updates() {
	new_version := strings.ToUpper(GetLatestVersion())
	version := strings.Split(new_version, " ")[1]
	if version != CheckSum() {
		color.White("[" + time.Now().Format("15:04:05.000000") + "] " + "DOWNLOADING " + new_version)
		if !DowloadUpdate() {
			color.Red("[" + time.Now().Format("15:04:05.000000") + "] " + "ERROR DOWNLOADING UPDATE")
			time.Sleep(2 * time.Second)
			os.Exit(255)
		}
		color.Yellow("[" + time.Now().Format("15:04:05.000000") + "] " + "UPDATE DOWNLOADED")
		time.Sleep(2 * time.Second)
		os.Exit(255)
	} else {
		color.Green("[" + time.Now().Format("15:04:05.000000") + "] " + "NO UPDATES")
		Version = version
	}

}
