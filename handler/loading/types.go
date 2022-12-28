package loading

import (
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
	Profiles []profile.Profile `json:"profiles"`
}
type Proxies struct {
	Proxies []settings.Proxie `json:"proxies"`
}

type Tasks struct {
	Tasks map[string][]string
}

type Config struct {
	// Accounts        Accounts
	// QuicktaskGroups map[int][]QuicktaskGroup
	Tasks    Tasks
	Proxies  Proxies
	Env      Env      `json:"env"`
	Settings Settings `json:"settings"`
	Profiles Profiles `json:"profiles"`
}
