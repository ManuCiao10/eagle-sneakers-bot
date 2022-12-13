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
	Profiles []profile.Profile
}

type Config struct {
	// Accounts        Accounts
	// Proxies         Proxies
	// Tasks           Tasks
	// QuicktaskGroups map[int][]QuicktaskGroup
	Env      Env      `json:"env"`
	Settings Settings `json:"settings"`
	Profiles Profiles
}
