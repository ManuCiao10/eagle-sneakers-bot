package loading

import (
	"github.com/eagle/eaglebot/handler/settings"
)

type Settings struct {
	Settings settings.Settings `json:"settings"`
}

type Config struct {
	// Accounts        Accounts
	// Proxies         Proxies
	// Tasks           Tasks
	// Profiles        Profiles
	// QuicktaskGroups map[int][]QuicktaskGroup
	Settings Settings `json:"settings"`
}