package thebrokenarm

import (
	"math/rand"
	"os"
	"strings"

	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/profile"
	"github.com/eagle/handler/settings"
	"github.com/eagle/handler/task"
	"github.com/fatih/color"
)

func GetProfile(t *task.Task) profile.Profile {
	for _, p := range loading.Data.Profiles.Profiles {
		if p.ID == t.Profile {
			return p
		}
	}

	return profile.Profile{
		ID: "not_found",
	}
}

func GetProxyList(t *task.Task) settings.Proxie {

	for _, proxy := range loading.Data.Proxies.Proxies {
		if proxy.ID == t.Proxy_List {
			return proxy
		}
	}

	return settings.Proxie{
		ID: "not_found",
	}
}

func RandomSize() string {
	sizes := []string{"36", "36.5", "37", "38"}

	return sizes[rand.Intn(len(sizes))]
}

func SplitSize(size string) string {
	sizes := strings.Split(size, ";")

	return sizes[rand.Intn(len(sizes))]
}

func err_(msg string) {
	color.Red(msg)
	os.Exit(0)
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
