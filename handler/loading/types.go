package loading

import (
	"github.com/eagle/handler/settings"
)

type Settings struct {
	Settings settings.Settings `json:"settings"`
}

type Env struct {
	Env settings.Env `json:"env"`
}

type Config struct {
	// Accounts        Accounts
	// Proxies         Proxies
	// Tasks           Tasks
	// Profiles        Profiles
	// QuicktaskGroups map[int][]QuicktaskGroup
	// Version  Version  `json:"version"`
	Env      Env      `json:"env"`
	Settings Settings `json:"settings"`
}
