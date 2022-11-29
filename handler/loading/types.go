package loading

import (
	"github.com/eagle/handler/settings"
	// "github.com/eagle/handler/version"
)

type Settings struct {
	Settings settings.Settings `json:"settings"`
}

// type Version struct {
// 	Version version.Info `json:"version"`
// }


type Config struct {
	// Accounts        Accounts
	// Proxies         Proxies
	// Tasks           Tasks
	// Profiles        Profiles
	// QuicktaskGroups map[int][]QuicktaskGroup
	// Version  Version  `json:"version"`
	Settings Settings `json:"settings"`
}
