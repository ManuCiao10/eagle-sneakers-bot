package thebrokenarm_monitor

import (
	"math/rand"
	"time"

	"github.com/eagle/handler/account"
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/profile"
	"github.com/eagle/handler/quicktask"
	"github.com/eagle/handler/settings"
)

func GetProfile(t *quicktask.Quicktask) profile.Profile {
	for _, p := range loading.Data.Profiles.Profiles {
		if p.ID == t.Profiles {
			return p
		}
	}

	return profile.Profile{
		ID: "not_found",
	}
}

func GetProxyList(t *quicktask.Quicktask) settings.Proxie {
	for _, proxy := range loading.Data.Proxies.Proxies {
		if proxy.ID == t.Proxylist {
			return proxy
		}
	}

	return settings.Proxie{
		ID: "not_found",
	}
}

// GetRandomAccount
func GetRandomAccount(siteId int) account.Account {
	accounts := loading.Data.Accounts.Accounts[siteId]

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(accounts))

	//check if the account is used
	if !accounts[randomIndex].Used {
		//if not used, set the account to used and return the account
		accounts[randomIndex].Used = true
		return accounts[randomIndex]
	} else {
		return account.Account{}
	}
}
