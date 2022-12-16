package thebrokenarm

import (
	"fmt"
	"os"

	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/profile"
	"github.com/eagle/handler/settings"
	"github.com/fatih/color"
)

func GetProfile(t *Task) profile.Profile {
	for _, p := range loading.Data.Profiles.Profiles {
		if p.ID == t.Profile {
			return p
		}
	}
	return profile.Profile{
		ID: "not_found",
	}
}

func GetProxyList(t *Task) settings.Proxie {

	for _, proxy := range loading.Data.Proxies.Proxies {
		fmt.Print(proxy.ID)
			fmt.Print(t.Proxy_List)
		if proxy.ID == t.Proxy_List {
			
			return proxy
		}
	}

	return settings.Proxie{
		ID: "not_found",
	}
}

func err_(msg string) {
	color.Red(msg)
	os.Exit(0)
}
