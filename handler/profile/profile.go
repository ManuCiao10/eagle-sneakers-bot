package profile

import (
	"sync"
)

var (
	profileMutex = sync.RWMutex{}
	profiles     = make(map[string]*Profile)
)

// DoesProfileExist checks if a profile exists
func DoesProfileExist(ID string) bool {
	profileMutex.RLock()
	defer profileMutex.RUnlock()

	_, ok := profiles[ID]
	return ok
}
