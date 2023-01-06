package loading

import (
	"github.com/eagle/handler/account"
	"github.com/eagle/handler/profile"
	"github.com/eagle/handler/settings"
)

type Settings struct {
	Settings settings.Settings `json:"settings"`
}

type Env struct {
	Env settings.Env `json:"env"`
}

type Profiles struct {
	Profiles []profile.Profile
}
type Proxies struct {
	Proxies []settings.Proxie
}

type Tasks struct {
	Tasks map[string][]string
}

type Quicktask struct {
	Quicktask map[string][]string
}

type Accounts struct {
	Accounts map[int][]account.Account
}

type Config struct {
	Accounts  Accounts
	Quicktask Quicktask
	Tasks     Tasks
	Proxies   Proxies
	Env       Env      `json:"env"`
	Settings  Settings `json:"settings"`
	Profiles  Profiles
}
