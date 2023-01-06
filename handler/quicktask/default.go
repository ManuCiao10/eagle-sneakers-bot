package quicktask

import (
	"strings"
	"sync"

	"github.com/lithammer/shortuuid"
)

var (
	quicktaskMutex = sync.RWMutex{}
	quickTask      = make(map[string]*Quicktask)
)

func CreateQuicktask(Site, Tasks_Quantity, Profiles, Accounts, Email, Proxylist, Payment_Method, Credit_Card, Other string) string {
	quicktaskMutex.Lock()
	defer quicktaskMutex.Unlock()

	id := shortuuid.New()

	quickTask[id] = &Quicktask{
		Site:           strings.ToLower(Site),
		Tasks_Quantity: strings.ToLower(Tasks_Quantity),
		Profiles:       strings.ToLower(Profiles),
		Accounts:       strings.ToLower(Accounts),
		Email:          strings.ToLower(Email),
		Proxylist:      strings.ToLower(Proxylist),
		Payment_Method: strings.ToLower(Payment_Method),
		Credit_Card:    strings.ToLower(Credit_Card),
		Other:          strings.ToLower(Other),
	}

	return id
}
