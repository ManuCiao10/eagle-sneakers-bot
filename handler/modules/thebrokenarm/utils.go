package thebrokenarm

import (
	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/profile"
)

func GetProfile(t *Task) profile.Profile {
	for _, p := range loading.Data.Profiles.Profiles {
		if p.ID == t.Profile {
			return p
		}
	}
	return profile.Profile{}
}
