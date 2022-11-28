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
	Version = "0.0.17"
	File    Update
)

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
	for _, v := range File.Data {
		if strings.Contains(v.Version, "Update") {
			return v.Version
		}
	}
	return Version
}

func DowloadUpdate() {
	color.Yellow("[" + time.Now().Format("15:04:05.000000") + "] " + "DOWNLOADING UPDATE...")
}

func Updates() {
	last := strings.Split(Version, ".")
	new_version := strings.ToUpper(GetLatestVersion())
	new := strings.Split(new_version, ".")
	if last[2] < new[2] {
		color.Cyan("[" + time.Now().Format("15:04:05.000000") + "] " + "NEW VERSION AVAILABLE " + new_version)
		DowloadUpdate()
		
	}
	time.Sleep(30 * time.Second)

}
