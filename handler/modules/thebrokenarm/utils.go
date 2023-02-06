package thebrokenarm

import (
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/profile"
	"github.com/eagle/handler/settings"
	"github.com/eagle/handler/task"
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
		if proxy.ID == t.CheckoutProxy {
			return proxy
		}
	}

	return settings.Proxie{
		ID: "not_found",
	}
}
