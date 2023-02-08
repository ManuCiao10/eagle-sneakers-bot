package version

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/logs"
)

var (
	Version = "0.0.23"
	File    Update
)

func GetLatestVersion() string {
	url := "https://api.hyper.co/v6/products"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+loading.Data.Env.Env.API_TOKEN)

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
		if strings.Contains(v.Version, "EagleBot") {
			return v.Files[0]
		}
	}
	return ""
}

func DowloadUpdate(version string) bool {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://manuciao5388.hyper.co/ajax/files/"+GetID(), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("cookie", "authorization="+loading.Data.Env.Env.AUTH_DOWNLOAD)
	req.Header.Set("hyper-account", loading.Data.Env.Env.ACC_DOWLOAD)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("EagleBot_" + version + ".exe")
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

func Initialize() {
	new_version := strings.ToUpper(GetLatestVersion())
	version := strings.Split(new_version, " ")[1]
	if version != Version {
		if !DowloadUpdate(version) {
			logs.LogsMsgErr("error downloading update")
		}
		logs.LogsMsgInfo("downloading... " + new_version)
		logs.LogsMsgSuccess("update downloaded")
	}
	Version = version

}

func GetVersion() string {
	return "[Eagle " + Version + "] "
}
