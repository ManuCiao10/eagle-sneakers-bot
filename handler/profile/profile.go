package profile

import (
	"errors"
	"sync"
)

var (
	profileMutex = sync.RWMutex{}

	ProfileDoesNotExistErr = errors.New("profile does not exist")
	ProfileNotAssignedErr  = errors.New("profile not assigned")
	profiles               = make(map[string]*Profile)
)

// CreateProfile creates a new profile
// func CreateProfile(profile *Profile) string {
// 	profileMutex.RLock()
// 	defer profileMutex.RUnlock()

// 	id := shortuuid.New()

// 	profiles[id] = profile

// 	return id
// }
